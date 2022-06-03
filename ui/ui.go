package ui

import (
	"fmt"
	"strings"

	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
)

type UI struct {
	lolClient  *lolclient.LoLClient
	settingsDB *settingsdb.SettingsDB
}

func NewUI(lolClient *lolclient.LoLClient, settingsDB *settingsdb.SettingsDB) *UI {
	return &UI{
		lolClient:  lolClient,
		settingsDB: settingsDB,
	}
}

func (ui *UI) Loop() error {
	fmt.Printf("Kisa is running. Enter a champion name to save current keybindings for that champion. Enter \"%s\" to save settings for all other champions.\n", settingsdb.DefaultName)
	for {
		var text string
		_, err := fmt.Scanln(&text)
		if err != nil {
			if err.Error() == "unexpected newline" {
				fmt.Printf("Empty input. Please try again\n")
				continue
			}

			return err
		}

		text = strings.Replace(text, "\"", "", -1)

		keyBindings, err := ui.lolClient.GetKeyBindings()
		if err != nil {
			return err
		}

		err = ui.settingsDB.PutKeyBindings(strings.ToLower(text), keyBindings)
		if err != nil {
			return err
		}

		fmt.Println("Key bindings saved. Please enter another champion name, or \"default\".")
	}
}
