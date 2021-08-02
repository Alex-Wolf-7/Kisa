package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const API_KEY = "RGAPI-641a0b82-0fbe-4c09-a6e0-bc756cf92521"
const SUMMONER = "Wolfee77"

const ENDPOINT_SUMMONER_V4 = "lol/summoner/v4/summoners/by-name/" // {/summoner_name}
const URL_NA = "https://na1.api.riotgames.com/"

func main() {
	client := http.DefaultClient

	req, err := http.NewRequest("GET", URL_NA+ENDPOINT_SUMMONER_V4+SUMMONER, nil)
	if err != nil {
		log.Fatalf("Unable to create new HTTP request: %s", err)
	}

	req.Header.Add("X-Riot-Token", API_KEY)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to perform HTTP request: %s", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)

	var summonerV4Resp models.SummonerV4Response
	if err := json.NewDecoder(resp.Body).Decode(&summonerV4Resp); err != nil {
		log.Fatalf("Unable to decode SummonerV4 response: %s", err)
	}

	log.Printf("Body: %s\n", summonerV4Resp)

	return
}
