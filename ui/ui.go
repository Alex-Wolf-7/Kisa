package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alex-Wolf-7/Kisa/gamesettings"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
)

type State int

type UI struct {
	reader     *bufio.Reader
	lolClient  *lolclient.LoLClient
	settingsDB *settingsdb.SettingsDB
	stop       bool
}

func NewUI(lolClient *lolclient.LoLClient, settingsDB *settingsdb.SettingsDB) (*UI, error) {
	reader := bufio.NewReader(os.Stdin)
	ui := &UI{
		reader:     reader,
		lolClient:  lolClient,
		settingsDB: settingsDB,
		stop:       false,
	}

	err := ui.checkDefault()
	if err != nil {
		return nil, fmt.Errorf("unable to check default settings: %s", err.Error())
	}

	return ui, nil
}

func (ui *UI) checkDefault() error {
	defaultSettings, err := ui.settingsDB.GetDefaultSettings()
	if err != nil {
		return fmt.Errorf("unable to get default settings: %s", err.Error())
	}

	if defaultSettings != nil {
		return nil
	}

	// No default settings: prompt for some and save
	fmt.Println("No default settings detected. Change your settings to a good default, then press \"Enter\"")
	ui.reader.ReadString('\n')
	defaultSettings, err = ui.lolClient.GetGameSettings()
	if err != nil {
		return fmt.Errorf("unable to get default settings: %s", err.Error())
	}
	err = ui.settingsDB.PutDefaultSettings(defaultSettings)
	if err != nil {
		return fmt.Errorf("unable to put default settings: %s", err.Error())
	}

	return nil
}

func (ui *UI) Loop() (error, bool) {
	for !ui.stop {
		// Get input
		fmt.Println("Please enter a champion name to save their game settings, or \"exit\" to quit. Enter \"default\" to change default settings")
		text := ui.readString()
		if ui.stop {
			return nil, false
		}
		if text == "exit" {
			plog.Infof("Goodbye!\n")
			return nil, true
		}
		if text == "" {
			continue
		}

		// Check current settings of named champion
		retrievedGameSettings, err := ui.settingsDB.GetSettings(text)
		if err != nil {
			return fmt.Errorf("unable to get game settings: %s", err.Error()), false

		} else if retrievedGameSettings != nil {
			ui.overwriteChampionSettings(text, retrievedGameSettings)
		} else {
			ui.saveNewChampionSettings(text)
		}

		fmt.Println() // Newline
	}

	return nil, false
}

// Champion does not exist: create new entry and compare vs default
func (ui *UI) saveNewChampionSettings(champion string) error {
	newSettings, err := ui.lolClient.GetGameSettings()
	if err != nil {
		return fmt.Errorf("Unable to get new game settings: %s", err.Error())
	}

	ui.settingsDB.PutSettings(champion, newSettings)
	fmt.Println("Champion settings saved")

	defaultSettings, err := ui.settingsDB.GetDefaultSettings()
	if err != nil {
		return fmt.Errorf("Unable to get default game settings: %s", err.Error())
	}

	settingsDiff := defaultSettings.GetChanges(newSettings)
	fmt.Println("Differences from default settings:")
	for outerKey, outerMap := range settingsDiff {
		for innerKey := range outerMap {
			propertyName := strings.TrimPrefix(innerKey, "evt")
			oldVal := defaultSettings[outerKey][innerKey]
			newVal := newSettings[outerKey][innerKey]
			fmt.Printf("%s: %s -> %s\n", propertyName, oldVal, newVal)
		}
	}

	return nil
}

// Champion exists: overwrite
func (ui *UI) overwriteChampionSettings(champion string, oldSettings gamesettings.GameSettings) error {
	newSettings, err := ui.lolClient.GetGameSettings()
	if err != nil {
		return fmt.Errorf("Unable to get new game settings: %s", err.Error())
	}

	ui.settingsDB.PutSettings(champion, newSettings)
	fmt.Println("Champion settings overwritten")

	settingsDiff := oldSettings.GetChanges(newSettings)
	fmt.Println("Differences from previously saved settings:")
	for outerKey, outerMap := range settingsDiff {
		for innerKey := range outerMap {
			propertyName := strings.TrimPrefix(innerKey, "evt")
			oldVal := oldSettings[outerKey][innerKey]
			newVal := newSettings[outerKey][innerKey]
			fmt.Printf("%s: %s -> %s\n", propertyName, oldVal, newVal)
		}
	}

	return nil
}

func (ui *UI) readString() string {
	text, _ := ui.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r\n", "", -1)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return text
}

func (ui *UI) Stop() {
	ui.stop = true
}
