package riot

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/Alex-Wolf-7/Kisa/match"
// 	"github.com/Alex-Wolf-7/Kisa/riot/interfaces"
// 	"github.com/Alex-Wolf-7/Kisa/riot/models"
// 	"github.com/Alex-Wolf-7/Kisa/summoner"
// )

// const (
// 	NA_SUMMONER_BY_NAME          = "lol/summoner/v4/summoners/by-name/%s" // summoner name
// 	AM_ENDPOINT_MATCHES_BY_PUUID = "lol/match/v5/matches/by-puuid/%s/ids" // puuid
// 	AM_MATCHES_BY_ID             = "lol/match/v5/matches/%s"              // matchID

// 	URL_NA       = "https://na1.api.riotgames.com/"
// 	URL_AMERICAS = "https://americas.api.riotgames.com/"

// 	HEADER_API_KEY = "X-Riot-Token"
// )

// type Client struct {
// 	apiKey     string
// 	httpClient *http.Client
// }

// func NewClient(apiKey string, httpClient *http.Client) interfaces.RiotClient {
// 	return &Client{
// 		apiKey:     apiKey,
// 		httpClient: httpClient,
// 	}
// }

// func (rc *Client) GetSummonerByName(name string) (*summoner.Summoner, error) {
// 	url := URL_NA + fmt.Sprintf(NA_SUMMONER_BY_NAME, name)
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to create GetSummonerByName request: %w", err)
// 	}

// 	req.Header.Add(HEADER_API_KEY, rc.apiKey)

// 	resp, err := rc.httpClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error in GetSummonerByName request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	var getSummonerByNameResp models.GetSummonerByNameResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&getSummonerByNameResp); err != nil {
// 		return nil, fmt.Errorf("Unable to decode GetSummonerByName response: %w", err)
// 	}

// 	revisionDate := time.Unix(0, getSummonerByNameResp.RevisionDate*int64(time.Millisecond))
// 	summoner := summoner.NewSummoner(getSummonerByNameResp.ID, getSummonerByNameResp.AccountID, getSummonerByNameResp.PUUID, getSummonerByNameResp.Name, getSummonerByNameResp.ProfileIconID, revisionDate, getSummonerByNameResp.SummonerLevel)

// 	return summoner, nil
// }

// func (rc *Client) GetMatchesByPUUID(puuid string, numMatches int) (models.Matches, error) {
// 	url := URL_AMERICAS + fmt.Sprintf(AM_ENDPOINT_MATCHES_BY_PUUID, puuid)
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to create GetMatchesByPUUID request: %w", err)
// 	}

// 	query := req.URL.Query()
// 	query.Add("count", strconv.Itoa(numMatches))
// 	req.URL.RawQuery = query.Encode()

// 	req.Header.Add(HEADER_API_KEY, rc.apiKey)
// 	resp, err := rc.httpClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error in GetMatchesByPUUID request: %w", err)
// 	}

// 	var matches []string
// 	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
// 		return nil, fmt.Errorf("Unable to decode GetMatchesByPUUID response: %w", err)
// 	}

// 	return matches, nil
// }

// func (rc *Client) GetMatchByID(matchID string) (*match.Match, error) {
// 	url := URL_AMERICAS + fmt.Sprintf(AM_MATCHES_BY_ID, matchID)
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to create GetMatchByID request: %w", err)
// 	}

// 	req.Header.Add(HEADER_API_KEY, rc.apiKey)
// 	resp, err := rc.httpClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error in GetMatchByID request: %w", err)
// 	}

// 	var match match.Match
// 	if err := json.NewDecoder(resp.Body).Decode(&match); err != nil {
// 		return nil, fmt.Errorf("Unable to decode GetMatchByID response: %w", err)
// 	}

// 	return &match, nil
// }
