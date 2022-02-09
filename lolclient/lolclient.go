package lolclient

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Alex-Wolf-7/Kisa/game"
	gs "github.com/Alex-Wolf-7/Kisa/gamesettings"
	"github.com/Alex-Wolf-7/Kisa/match"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/summoner"
)

const (
	CURRENT_SUMMONER     = "lol-summoner/v1/current-summoner"
	CHAMP_SELECT_SESSION = "lol-champ-select/v1/session"
	GAME_SESSION         = "lol-gameflow/v1/session"
	GAME_SETTINGS        = "lol-game-settings/v1/input-settings"
)

type LoLClient struct {
	http      *http.Client
	authToken string
	url       string
}

func NewLoLClient(authToken string, clientURL string, port string) *LoLClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := http.DefaultClient
	client.Transport = &transport

	return &LoLClient{
		http:      client,
		authToken: authToken,
		url:       fmt.Sprintf(clientURL, port),
	}
}

func (lol *LoLClient) GetCurrentSummoner() (*summoner.Summoner, error) {
	reqUrl := lol.url + CURRENT_SUMMONER
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating GetCurrentSummoner request: %s", err)
	}
	lol.setAuthorizationHeader(req)

	resp, err := lol.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error performing GetCurrentSummoner request: %s", err)
	}

	summoner := new(summoner.Summoner)
	err = json.NewDecoder(resp.Body).Decode(summoner)
	if err != nil {
		return nil, fmt.Errorf("Error decoding Summoner object from GetCurrentSummoner response: %s", err)
	}

	return summoner, nil
}

func (lol *LoLClient) GetChampSelectSession() (*match.Match, error) {
	reqUrl := lol.url + CHAMP_SELECT_SESSION
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating GetChampSelectSession request: %s", err)
	}
	lol.setAuthorizationHeader(req)

	resp, err := lol.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error performing GetChampSelectSession request: %s", err)
	}

	match := new(match.Match)
	err = json.NewDecoder(resp.Body).Decode(match)
	if err != nil {
		return nil, fmt.Errorf("Error decoding Match object from GetChampSelectSession response: %s", err)
	}

	return match, nil
}

func (lol *LoLClient) GetGameSession() (*game.Game, error) {
	reqUrl := lol.url + GAME_SESSION
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating GetGameSession request: %s", err)
	}
	lol.setAuthorizationHeader(req)

	resp, err := lol.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error performing GetGameSession request: %s", err)
	}

	game := new(game.Game)
	err = json.NewDecoder(resp.Body).Decode(game)
	if err != nil {
		return nil, fmt.Errorf("Error decoding Game object from GetGameSession response: %s", err)
	}

	return game, nil
}

func (lol *LoLClient) GetGameSettings() (gs.GameSettings, error) {
	reqUrl := lol.url + GAME_SETTINGS
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating GetGameSettings request: %s", err)
	}
	lol.setAuthorizationHeader(req)

	resp, err := lol.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error performing GetGameSettings request: %s", err)
	}

	gameSettings := gs.NewGameSettings()
	err = json.NewDecoder(resp.Body).Decode(&gameSettings)
	if err != nil {
		return nil, fmt.Errorf("Error decoding GameSettings object from GetGameSettings response: %s", err)
	}

	return gameSettings, nil
}

func (lol *LoLClient) PatchGameSettings(gameSettings gs.GameSettings) error {
	settingsBytes, err := json.Marshal(gameSettings)
	if err != nil {
		return fmt.Errorf("Unable to marshal gameSettings into JSON: %s", err)
	}

	reqBody := bytes.NewReader(settingsBytes)

	reqUrl := lol.url + GAME_SETTINGS
	req, err := http.NewRequest("PATCH", reqUrl, reqBody)
	if err != nil {
		return fmt.Errorf("Error creating PatchGameSettings request: %s", err)
	}
	lol.setAuthorizationHeader(req)

	resp, err := lol.http.Do(req)
	if err != nil {
		return fmt.Errorf("Error performing GetGameSettings request: %s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("PatchGameSettings response code is not 200: %s", resp.Status)
	}

	return nil
}

// Patches game settings, then checks to make sure it has been patched correctly
func (lol *LoLClient) PatchGameSettingsMultiple(gameSettings gs.GameSettings) error {
	startGameSettings, err := lol.GetGameSettings()
	if err != nil {
		return err
	}
	startGameSettingsJson, err := json.Marshal(startGameSettings)
	if err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		err = lol.PatchGameSettings(gameSettings)
		if err != nil {
			return err
		}

		afterGameSettings, err := lol.GetGameSettings()
		if err != nil {
			return err
		}
		afterGameSettingsJson, err := json.Marshal(afterGameSettings)
		if err != nil {
			return err
		}

		if string(startGameSettingsJson) != string(afterGameSettingsJson) {
			plog.Debugf("Successfully patched game settings. Took %d tries", i+1)
			return nil
		} else {
			plog.Debugf("Unable to patch game settings, trying again")
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func (lol *LoLClient) setAuthorizationHeader(req *http.Request) {
	req.Header.Add("Authorization", lol.buildAuthorizationToken())
}

func (lol *LoLClient) buildAuthorizationToken() string {
	preEncode := "riot:" + lol.authToken
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(preEncode))
}
