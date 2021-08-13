package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Alex-Wolf-7/Kisa/lolclient"
)

type State int

const (
	State_START = iota
)

type IO struct {
	state               State
	reader              *bufio.Reader
	lolClient           *lolclient.LoLClient
	allGameSettings     map[string]map[string]map[string]interface{}
	defaultGameSettings map[string]map[string]interface{}
}

func Start(lolClient *lolclient.LoLClient) int {
	reader := bufio.NewReader(os.Stdin)
	io := IO{
		state:               State_START,
		reader:              reader,
		lolClient:           lolClient,
		allGameSettings:     make(map[string]map[string]map[string]interface{}),
		defaultGameSettings: make(map[string]map[string]interface{}),
	}

	fmt.Println("Grabbing default settings...")
	// defaultGameSettings, err := lolClient.GetGameSettings()
	defaultGameSettings, err := lolClient.GetGameSettings()
	if err != nil {
		log.Printf("Error getting default game settings from League Client: %s\n", err)
		return 1
	}
	io.defaultGameSettings = defaultGameSettings

	for {
		fmt.Println("Please enter a champion name to save their game settings, or \"exit\" to quit")
		text, _ := io.reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r\n", "", -1)
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)
		if text == "exit" {
			log.Println("Goodbye")
			return 0
		}

		retrievedGameSettings, ok := io.allGameSettings[text]
		if ok {
			err = lolClient.PatchGameSettings(retrievedGameSettings)
			if err != nil {
				log.Printf("Error patching game settings: %s", err)
				return 1
			}
			fmt.Printf("Changed your settings to %s!\n", text)

		} else {
			newGameSettings, err := lolClient.GetGameSettings()
			if err != nil {
				log.Printf("Error getting champion game settings from League Client: %s\n", err)
				return 1
			}

			io.allGameSettings[text] = newGameSettings
			settingsDiff := getSettingsDiff(io.defaultGameSettings, newGameSettings)

			fmt.Println("Champion settings saved")
			fmt.Println("Differences:")
			for outerKey, outerMap := range settingsDiff {
				for innerKey := range outerMap {
					propertyName := strings.TrimPrefix(innerKey, "evt")
					oldVal := io.defaultGameSettings[outerKey][innerKey]
					newVal := newGameSettings[outerKey][innerKey]
					fmt.Printf("%s: %s -> %s\n", propertyName, oldVal, newVal)
				}
			}
		}

		fmt.Println()
	}
}

func getSettingsDiff(old map[string]map[string]interface{}, new map[string]map[string]interface{}) map[string]map[string]bool {
	addMap := make(map[string]map[string]bool)
	for outerKey, oldInnerMap := range old {
		for innerKey, oldValue := range oldInnerMap {
			newValue, ok := new[outerKey][innerKey]
			if !ok || newValue != oldValue {
				addMap[outerKey] = safeMapBoolAdd(addMap[outerKey], innerKey, true)
			}
		}
	}

	for outerKey, newInnerMap := range new {
		for innerKey, newValue := range newInnerMap {
			oldValue, ok := old[outerKey][innerKey]
			if !ok || newValue != oldValue {
				addMap[outerKey] = safeMapBoolAdd(addMap[outerKey], innerKey, true)
			}
		}
	}

	return addMap
}

func safeMapBoolAdd(m map[string]bool, k string, v bool) map[string]bool {
	if m == nil {
		m = make(map[string]bool)
	}

	m[k] = v
	return m
}
