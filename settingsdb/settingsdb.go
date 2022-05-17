package settingsdb

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/gamesettings"
	"github.com/Alex-Wolf-7/Kisa/keybindings"
	"github.com/Alex-Wolf-7/Kisa/opsys"
)

const (
	DefaultName = "default"

	directoryPerms = 0755
	filePerms      = 0644
)

type SettingsDB struct {
	path  string
	opSys opsys.OpSys
}

func NewSettingsDB(opSys opsys.OpSys) (*SettingsDB, error) {
	var dbPath string
	if opSys.IsMac() {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("unable to find home directory: %s", err.Error())
		}
		dbPath = fmt.Sprintf(constants.MAC_SETTINGS_DB_PATH, homeDir)
	} else if opSys.IsWindows() {
		userName := os.Getenv("USERNAME")
		dbPath = fmt.Sprintf(constants.WINDOWS_SETTINGS_DB_PATH, userName)
	} else {
		return nil, fmt.Errorf("unrecognized OS: %s", opSys.String())
	}

	// If data directory does not exist, make it
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, directoryPerms)
		if err != nil {
			return nil, fmt.Errorf("unable to build settings db. Path not valid: %s", err.Error())
		}
	}

	fmt.Println(dbPath)
	err := os.Chdir(dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to change into settings directory: %s", err.Error())
	}

	return &SettingsDB{
		opSys: opSys,
		path:  dbPath,
	}, nil
}

func (db *SettingsDB) PutKeyBindings(champion string, keyBindings *keybindings.KeyBindings) error {
	fileName := championToFileName(champion)

	data, err := keyBindings.MarshalJSON(true)
	if err != nil {
		return fmt.Errorf("unable to marshal settings into json: %s", err.Error())
	}

	os.WriteFile(fileName, data, filePerms)
	return nil
}

func (db *SettingsDB) PutDefaultKeyBindings(keyBindings *keybindings.KeyBindings) error {
	return db.PutKeyBindings(DefaultName, keyBindings)
}

func (db *SettingsDB) GetKeyBindings(champion string) (*keybindings.KeyBindings, error) {
	fileName := championToFileName(champion)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Does not exist
		return nil, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read champion settings data file: %s", err.Error())
	}

	if len(data) == 0 {
		return nil, nil
	}

	keyBindings := new(keybindings.KeyBindings)
	err = json.Unmarshal(data, keyBindings)
	if err != nil {
		return nil, fmt.Errorf("unable parse data file into champion key bindings: %s", err.Error())
	}

	return keyBindings, nil
}

func (db *SettingsDB) GetDefaultKeyBindings() (*keybindings.KeyBindings, error) {
	return db.GetKeyBindings(DefaultName)
}

func (db *SettingsDB) PutSettings(champion string, settings gamesettings.GameSettings) error {
	fileName := championToFileName(champion)

	data, err := json.Marshal(settings)
	if err != nil {
		return fmt.Errorf("unable to marshal settings into json: %s", err.Error())
	}

	os.WriteFile(fileName, data, filePerms)
	return nil
}

func (db *SettingsDB) PutDefaultSettings(settings gamesettings.GameSettings) error {
	return db.PutSettings(DefaultName, settings)
}

func (db *SettingsDB) GetSettings(champion string) (gamesettings.GameSettings, error) {
	fileName := championToFileName(champion)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Does not exist
		return nil, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read champion settings data file: %s", err.Error())
	}

	if len(data) == 0 {
		return nil, nil
	}

	var gameSettings gamesettings.GameSettings
	err = json.Unmarshal(data, &gameSettings)
	if err != nil {
		return nil, fmt.Errorf("unable parse data file into champion settings: %s", err.Error())
	}

	return gameSettings, nil
}

func (db *SettingsDB) GetAllSettings() {}

func (db *SettingsDB) GetDefaultSettings() (gamesettings.GameSettings, error) {
	return db.GetSettings(DefaultName)
}

func championToFileName(champion string) string {
	return strings.ToLower(champion) + ".json"
}
