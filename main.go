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
	iopublisher "github.com/Alex-Wolf-7/Kisa/iopublisher"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/opsys"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
	"github.com/Alex-Wolf-7/Kisa/ui"
)

var clientErrorDisplayed = false

func main() {
	opSys := opsys.NewOpSys(runtime.GOOS)
	settingsDB, err := settingsdb.NewSettingsDB(opSys)
	if err != nil {
		plog.FatalfWithCode("100", "unable to set up settings db: %s\n", err.Error())
	}

	ioPublisher := iopublisher.NewIOPublisher()

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
		} else if summoner == nil || summoner.InternalName == "" {
			plog.ErrorfWithBackup("Unable to get current summoner: Restarting\n", "unable to get current summoner: empty\n")
		} else {
			plog.Infof("Summoner: %s\n", summoner.InternalName)
		}

		b := background.NewBackground(lolClient, settingsDB, summoner, championMap)

		ui, err := ui.NewUI(lolClient, settingsDB)
		if err != nil {
			plog.ErrorfWithBackup("Unable to start app: Restarting\n", "unable to open UI: %s\n", err.Error())
			continue
		}

		// Starts UI thread in different channel
		ioPublisher.Listen(ui)

		// Sustained background loop, blocks until error
		err = b.Loop()

		// Cleanup
		ioPublisher.RemoveListener()
		if err != nil {
			plog.ErrorfWithBackup("Error: restarting\n", "Error in background thread: %s\n", err.Error())
		}
		plog.Debugf("Background thread killed\n")
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
		if !clientErrorDisplayed {
			fmt.Println("Please open the League of Legends client")
			clientErrorDisplayed = true
		} else {
			plog.Periodicf("League of Legends client must be running (portMatch)\n")
		}

		return "", "", nil
	} else {
		// Show "please open client" message only once until client actually opens
		clientErrorDisplayed = false
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

func runBackground(background *background.Background, ui *ui.UI, failChan chan bool, ioPublisher *iopublisher.IOPublisher) {
	go func() {
		err := background.Loop()
		ioPublisher.RemoveListener()
		failChan <- true
		if err != nil {
			plog.ErrorfWithBackup("Error: restarting\n", "Error in background thread: %s\n", err.Error())
		}
		plog.Debugf("Background thread killed\n")
	}()
}
