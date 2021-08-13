package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/Alex-Wolf-7/Kisa/game"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/summoner"
	"github.com/Alex-Wolf-7/Kisa/ui"
)

const (
	API_KEY    = "RGAPI-641a0b82-0fbe-4c09-a6e0-bc756cf92521"
	URL_CLIENT = "https://127.0.0.1:%s/" // Port number
)

func main() {
	var port, authToken string
	if runtime.GOOS == "windows" {
		port, authToken = getClientInfoWindows()
	} else {
		port, authToken = getClientInfoMac()
	}

	lolClient := lolclient.NewLoLClient(authToken, URL_CLIENT, port)

	code := ui.Start(lolClient)
	os.Exit(code)
}

func main2() {
	var port, authToken string
	if runtime.GOOS == "windows" {
		port, authToken = getClientInfoWindows()
	} else {
		port, authToken = getClientInfoMac()
	}

	lolClient := lolclient.NewLoLClient(authToken, URL_CLIENT, port)
	gameSettings, err := lolClient.GetGameSettings()
	if err != nil {
		log.Fatalln(err)
	}

	fullGameSettings := make(map[string]*gamesettings.GameSettings)
	fullGameSettings["default"] = gameSettings

	// summoner, err := lolClient.GetCurrentSummoner()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// dataDragon := datadragon.NewDataDragonClient()
	// clientVersion, err := dataDragon.GetMostRecentVersion()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// championByKey, err := dataDragon.GetChampionMapByKey(clientVersion)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// champID, err := getChampWhenLockedIn(lolClient, summoner)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// champ, ok := championByKey[champID]
	// if !ok {
	// 	log.Fatalf("Champion ID %s does not exist in champion map\n", champID)
	// }
	// fmt.Println("Champ:", champ.Name)
}

func getClientInfoWindows() (string, string) {
	log.Fatalln("Not implemented for Windows yet")
	return "bleep", "bloop"
}

func getClientInfoMac() (string, string) {
	cmd := exec.Command("bash", "-c", "ps -A | grep LeagueClientUx")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Unable to get current mac processes: %s\n", err)
	}

	portMatcher, err := regexp.Compile("--app-port=([0-9]*)")
	if err != nil {
		log.Fatalf("Unable to build Regex matcher for app port: %s\n", err)
	}

	portMatch := string(portMatcher.Find(out))
	if portMatch == "" {
		log.Fatalln("League of Legends client must be running")
	}
	port := strings.Split(portMatch, "=")[1]

	authMatcher, err := regexp.Compile("--remoting-auth-token=([0-9A-Za-z_-]*)")
	if err != nil {
		log.Fatalf("Unable to build Regex matcher for auth token: %s\n", err)
	}

	authMatch := string(authMatcher.Find(out))
	if authMatch == "" {
		log.Fatalln("League of Legends client must be running")
	}
	auth := strings.Split(authMatch, "=")[1]

	return port, auth
}

func getChampWhenLockedIn(lolClient *lolclient.LoLClient, summoner *summoner.Summoner) (string, error) {
	for {
		currentGame, err := lolClient.GetGameSession()
		if err != nil {
			return "", err
		}

		if currentGame.GetPhase() == game.Phase_IN_PROGRESS {
			champ := currentGame.GetChampionForSummoner(summoner.InternalName)
			return champ, nil
		} else if currentGame.IsValid() {
			fmt.Println("Game preparing:", currentGame.Phase)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("No game started")
			time.Sleep(10 * time.Second)
		}
	}
}
