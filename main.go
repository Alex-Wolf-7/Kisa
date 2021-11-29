package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/Alex-Wolf-7/Kisa/background"
	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/datadragon"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/opsys"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
	"github.com/Alex-Wolf-7/Kisa/ui"
)

func main() {
	opSys := opsys.NewOpSys(runtime.GOOS)
	settingsDB, err := settingsdb.NewSettingsDB(opSys)
	if err != nil {
		plog.FatalfWithCode("100", "unable to set up settings db: %s\n", err.Error())
	}

	for {
		port, authToken, err := waitForPortAndToken(opSys)
		if err != nil {
			plog.ErrorfWithBackup("Unable to get client data: retrying\n", "error getting port and auth token: %s\n", err.Error())
			continue
		}

		lolClient := lolclient.NewLoLClient(authToken, constants.URL_CLIENT_FORMAT, port)

		dataDragon := datadragon.NewDataDragonClient()
		version, err := dataDragon.GetMostRecentVersion()
		if err != nil {
			plog.ErrorfWithBackup("Unable to connect to Riot databases: restarting\n", "unable to get most recent datadragon version: %s\n", err.Error())
			continue
		}
		championMap, err := dataDragon.GetChampionMapByKey(version)
		if err != nil {
			plog.ErrorfWithBackup("Unable to get data from Riot: restarting\n", "unable to get champion map by key: %s\n", err.Error())
			continue
		}

		summoner, err := lolClient.GetCurrentSummoner()
		if err != nil {
			plog.ErrorfWithBackup("Unable to get current summoner: Restarting\n", "unable to get current summoner: %s\n", err.Error())
			continue
		}

		b := background.NewBackground(lolClient, settingsDB, summoner, championMap)

		ui, err := ui.NewUI(lolClient, settingsDB)
		if err != nil {
			plog.ErrorfWithBackup("Unable to start app: Restarting\n", "unable to open UI: %s\n", err.Error())
			continue
		}

		failChan := make(chan bool, 2)
		runUI(ui, b, failChan)
		runBackground(b, ui, failChan)
		<-failChan
	}
}

func waitForPortAndToken(opSys opsys.OpSys) (string, string, error) {
	var port, authToken string
	var err error

	for {
		port, authToken, err = getClientInfo(opSys)

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

func getProcessesWindows() ([]byte, error) {
	cmd := exec.Command("wmic", "PROCESS", "WHERE", "name='LeagueClientUx.exe'", "GET", "commandline")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Unable to get current windows processes: %s", err)
	}

	return out, nil
}

func getProcessesMac() ([]byte, error) {
	cmd := exec.Command("bash", "-c", "ps -A | grep LeagueClientUx")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Unable to get current mac processes: %s", err)
	}

	return out, nil
}

func getClientInfo(opSys opsys.OpSys) (string, string, error) {
	var out []byte
	var err error

	if opSys.IsWindows() {
		out, err = getProcessesWindows()
	} else if opSys.IsMac() {
		out, err = getProcessesMac()
	} else {
		return "", "", fmt.Errorf("Unrecognized GOOS: %s", opSys.String())
	}

	if err != nil {
		return "", "", err
	}

	portMatcher, err := regexp.Compile("--app-port=([0-9]*)")
	if err != nil {
		return "", "", fmt.Errorf("Unable to build Regex matcher for app port: %s", err)
	}

	portMatch := string(portMatcher.Find(out))
	if portMatch == "" {
		// No error; retry
		plog.Periodicf("League of Legends client must be running (portMatch)\n")
		return "", "", nil
	}
	port := strings.Split(portMatch, "=")[1]

	authMatcher, err := regexp.Compile("--remoting-auth-token=([0-9A-Za-z_-]*)")
	if err != nil {
		return "", "", fmt.Errorf("Unable to build Regex matcher for auth token: %s", err)
	}

	authMatch := string(authMatcher.Find(out))
	if authMatch == "" {
		// No error; retry
		plog.Periodicf("League of Legends client must be running (authMatch)\n")
		return "", "", nil
	}
	auth := strings.Split(authMatch, "=")[1]

	return port, auth, nil
}

func runBackground(background *background.Background, ui *ui.UI, failChan chan bool) {
	go func() {
		err := background.Loop()
		ui.Stop()
		failChan <- true
		if err != nil {
			plog.ErrorfWithBackup("Error: restarting\n", "Error in background thread: %s\n", err.Error())
		}
		plog.Debugf("Background thread killed\n")
	}()
}

func runUI(ui *ui.UI, background *background.Background, failChan chan bool) {
	go func() {
		err := ui.Loop()
		background.Stop()
		failChan <- true
		if err != nil {
			plog.ErrorfWithBackup("Error: restarting\n", "Error in UI thread: %s\n", err.Error())
		}
		plog.Debugf("UI thread killed\n")
	}()
}
