package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/opsys"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
	"github.com/Alex-Wolf-7/Kisa/ui"
)

func main() {
	opSys := opsys.NewOpSys(runtime.GOOS)
	settingsDB, err := settingsdb.NewSettingsDB(opSys)
	if err != nil {
		log.Fatalf("unable to set up settings db: %s", err.Error())
	}

	port, authToken, err := waitForPortAndToken(opSys)
	if err != nil {
		log.Fatalf("error getting port and auth token: %s", err.Error())
	}

	lolClient := lolclient.NewLoLClient(authToken, constants.URL_CLIENT_FORMAT, port)

	err = ui.Start(lolClient, settingsDB)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func waitForPortAndToken(opSys opsys.OpSys) (string, string, error) {
	var port, authToken string
	var err error

	for {
		if opSys.IsWindows() {
			port, authToken, err = getClientInfoWindows()
		} else if opSys.IsMac() {
			port, authToken, err = getClientInfoMac()
		} else {
			return "", "", fmt.Errorf("Unrecognized OpSys: %s", opSys.String())
		}

		if err != nil {
			return "", "", err
		} else if port != "" || authToken != "" {
			return port, authToken, nil
		} else {
			time.Sleep(constants.CHECK_IF_CLIENT_OPEN_TIME)
			continue
		}
	}
}

func getClientInfoWindows() (string, string, error) {
	return "", "", errors.New("Not implemented for Windows yet")
}

func getClientInfoMac() (string, string, error) {
	cmd := exec.Command("bash", "-c", "ps -A | grep LeagueClientUx")
	out, err := cmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("Unable to get current mac processes: %s\n", err)
	}

	portMatcher, err := regexp.Compile("--app-port=([0-9]*)")
	if err != nil {
		return "", "", fmt.Errorf("Unable to build Regex matcher for app port: %s\n", err)
	}

	portMatch := string(portMatcher.Find(out))
	if portMatch == "" {
		// No error; retry
		log.Print("League of Legends client must be running (portMatch)")
		return "", "", nil
	}
	port := strings.Split(portMatch, "=")[1]

	authMatcher, err := regexp.Compile("--remoting-auth-token=([0-9A-Za-z_-]*)")
	if err != nil {
		return "", "", fmt.Errorf("Unable to build Regex matcher for auth token: %s\n", err)
	}

	authMatch := string(authMatcher.Find(out))
	if authMatch == "" {
		// No error; retry
		log.Print("League of Legends client must be running (authMatch)")
		return "", "", nil
	}
	auth := strings.Split(authMatch, "=")[1]

	return port, auth, nil
}

// func getChampWhenLockedIn(lolClient *lolclient.LoLClient, summoner *summoner.Summoner) (string, error) {
// 	for {
// 		currentGame, err := lolClient.GetGameSession()
// 		if err != nil {
// 			return "", err
// 		}

// 		if currentGame.GetPhase() == game.Phase_IN_PROGRESS {
// 			champ := currentGame.GetChampionForSummoner(summoner.InternalName)
// 			return champ, nil
// 		} else if currentGame.IsValid() {
// 			fmt.Println("Game preparing:", currentGame.Phase)
// 			time.Sleep(2 * time.Second)
// 		} else {
// 			fmt.Println("No game started")
// 			time.Sleep(10 * time.Second)
// 		}
// 	}
// }
