package datadragon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alex-Wolf-7/Kisa/datadragon/models"
)

const (
	URL_DATADRAGON     = "https://ddragon.leagueoflegends.com/"
	ENDPOINT_VERSIONS  = "api/versions.json"
	ENDPOINT_CHAMPIONS = "cdn/%s/data/en_US/champion.json" // version
)

type Version string

type DataDragonClient struct {
	http *http.Client
}

func NewDataDragonClient() *DataDragonClient {
	http := http.DefaultClient

	return &DataDragonClient{
		http: http,
	}
}

func (ddc *DataDragonClient) GetVersions() ([]Version, error) {
	resp, err := ddc.http.Get(URL_DATADRAGON + ENDPOINT_VERSIONS)
	if err != nil {
		return nil, fmt.Errorf("Unable to get version list from DataDragon: %s", err)
	}

	var versions []Version
	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode versions from DataDragon response: %s", err)
	}
	if len(versions) == 0 {
		return nil, fmt.Errorf("No versions found")
	}

	return versions, nil
}

func (ddc *DataDragonClient) GetMostRecentVersion() (Version, error) {
	versions, err := ddc.GetVersions()
	if err != nil {
		return "", err
	}
	if len(versions) == 0 {
		return "", fmt.Errorf("No versions to choose most recent of, but also no error. This case should be handled in GetVersions and should not be reachable")
	}

	return versions[0], nil
}

func (ddc *DataDragonClient) GetUnformattedChampionList(clientVersion Version) (*models.ChampionListUnformatted, error) {
	resp, err := ddc.http.Get(URL_DATADRAGON + fmt.Sprintf(ENDPOINT_CHAMPIONS, clientVersion))
	if err != nil {
		return nil, fmt.Errorf("Unable to get champion list from DataDragon: %s", err)
	}

	champListNaive := new(models.ChampionListUnformatted)

	err = json.NewDecoder(resp.Body).Decode(champListNaive)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse champion list from GetChampionList response: %s", err)
	}

	return champListNaive, nil
}

func (ddc *DataDragonClient) GetChampionMapByKey(clientVersion Version) (map[string]*models.Champion, error) {
	champListNaive, err := ddc.GetUnformattedChampionList(clientVersion)
	if err != nil {
		return nil, err
	}

	championsByKey := make(map[string]*models.Champion)

	for _, obj := range champListNaive.Data {
		champion, err := models.BuildChampion(obj)
		if err != nil {
			return nil, err
		}
		championsByKey[champion.Key] = champion
	}

	return championsByKey, nil
}
