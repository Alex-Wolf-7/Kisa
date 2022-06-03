package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
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

var clientErrorDisplayed = false

func main() {
	opSys := opsys.NewOpSys(runtime.GOOS)
	settingsDB, err := settingsdb.NewSettingsDB(opSys)

	var startLock sync.Mutex

	go func() {
		for {
			b := setupBackground(opSys, settingsDB, &startLock)
			err = b.Loop()
			if err != nil {
				plog.ErrorfWithBackup("Background process quit\n", "background loop failed: %s\n", err.Error())
			}
		}
	}()

	for {
		ui := setupUI(opSys, settingsDB, &startLock)
		err = ui.Loop()
		if err != nil {
			plog.ErrorfWithBackup("Error with user interface\n", "ui loop failed: %s\n", err.Error())
		}
	}
}

func setupUI(opSys opsys.OpSys, settingsDB *settingsdb.SettingsDB, startLock *sync.Mutex) *ui.UI {
	port, authToken, err := waitForPortAndToken(opSys, startLock)
	if err != nil {
		plog.FatalfWithCode("001", "error getting port and auth token: %s\n", err.Error())
	}
	lolClient := lolclient.NewLoLClient(authToken, constants.URL_CLIENT_FORMAT, port)
	return ui.NewUI(lolClient, settingsDB)
}

func setupBackground(opSys opsys.OpSys, settingsDB *settingsdb.SettingsDB, startLock *sync.Mutex) *background.Background {
	port, authToken, err := waitForPortAndToken(opSys, startLock)
	if err != nil {
		plog.FatalfWithCode("001", "error getting port and auth token: %s\n", err.Error())
	}

	lolClient := lolclient.NewLoLClient(authToken, constants.URL_CLIENT_FORMAT, port)

	dataDragon := datadragon.NewDataDragonClient()
	version, err := dataDragon.GetMostRecentVersion()
	if err != nil {
		plog.FatalfWithCode("101", "unable to get most recent datadragon version: %s\n", err.Error())
	}
	championMap, err := dataDragon.GetChampionMapByKey(version)
	if err != nil {
		plog.FatalfWithCode("102", "unable to get champion map by key: %s\n", err.Error())
	}

	return background.NewBackground(lolClient, championMap, settingsDB)
}

func waitForPortAndToken(opSys opsys.OpSys, startLock *sync.Mutex) (string, string, error) {
	startLock.Lock()
	defer startLock.Unlock()

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
