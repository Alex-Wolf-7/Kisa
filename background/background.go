package background

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/datadragon/models"
	"github.com/Alex-Wolf-7/Kisa/game"
	"github.com/Alex-Wolf-7/Kisa/keybindings"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
)

type Background struct {
	lolClient    *lolclient.LoLClient
	championMap  map[string]*models.Champion
	settingsDB   *settingsdb.SettingsDB
	lastSetChamp string
}

func NewBackground(
	lolClient *lolclient.LoLClient,
	championMap map[string]*models.Champion,
	settingsDB *settingsdb.SettingsDB,
) *Background {

	return &Background{
		lolClient:   lolClient,
		championMap: championMap,
		settingsDB:  settingsDB,
	}
}

func (b *Background) Loop() error {
	plog.Infof("Background loop started")
	var champ *models.Champion
	gameStarted := false
	for {
		currentGame, err := b.lolClient.GetGameSession()
		if err != nil {
			return err
		}

		switch currentGame.GetPhase() {
		case game.Phase_CHAMP_SELECT:
			champ, err = b.ChampSelect()
			if err != nil {
				return err
			}
			time.Sleep(constants.CHECK_IF_GAME_STARTED_TIME)

		case game.Phase_IN_PROGRESS:
			gameStarted = true
			// Game started: be done
			continue
		default:
			if gameStarted {
				gameStarted = false
				// If game end, update keybindings for champion
				go func() {
					if champ != nil {
						time.Sleep(constants.AFTER_GAME_WAIT_TO_GET_KEYBINDINGS_TIME)
						keyBindings, err := b.lolClient.GetKeyBindings()
						if err != nil {
							plog.ErrorfWithBackup("Failed to set keybindings\n", "failed to get after-game keybindings for %s\n", champ.Name)
							return
						}
						err = b.settingsDB.PutKeyBindings(champ.Name, keyBindings)
						if err != nil {
							plog.ErrorfWithBackup("Failed to set keybindings\n", "failed to save after-game keybindings for %s\n", champ.Name)

						}
						fmt.Printf("Saved keybindings for %s\n", champ.Name)
						champ = nil
					}
				}()
			}
			// No champ select or game started: wait for game start
			time.Sleep(constants.CHECK_IF_CHAMP_SELECT_TIME)
			continue
		}
	}
}

func (b *Background) ChampSelect() (*models.Champion, error) {
	// In champ select: keep getting most up-to-date locked champion and set keybindings
	champNum, err := b.lolClient.GetLockedChampion()
	if err != nil {
		return nil, err
	}

	var champ *models.Champion
	if champNum != 0 {
		champ = b.championMap[strconv.Itoa(champNum)]

		keyBindings, err := b.settingsDB.GetKeyBindings(champ.Name)
		if err != nil {
			return nil, err
		}
		if keyBindings == nil {
			keyBindings, err = b.settingsDB.GetDefaultKeyBindings()
			if err != nil {
				return nil, err
			}
			if keyBindings == nil {
				keyBindings, err = b.NoDefaultKeybindings()
				if err != nil {
					return nil, err
				}
				fmt.Println("No default key bindings found. Set current key bindings as default.")
			}
		}

		err = b.lolClient.PatchKeyBindings(*keyBindings)
		if err != nil {
			plog.ErrorfWithBackup("unable to set keybindings", "unable to set keybindings to champion %s", champ.Name)
		}

		if champ.Name != b.lastSetChamp {
			fmt.Printf("Successfully set keybindings for %s\n", champ.Name)
			b.lastSetChamp = champ.Name
		}
	}
	return champ, nil
}

func (b *Background) NoDefaultKeybindings() (*keybindings.KeyBindings, error) {
	defaultBindings, err := b.lolClient.GetKeyBindings()
	if err != nil {
		return nil, err
	}

	err = b.settingsDB.PutDefaultKeyBindings(defaultBindings)
	if err != nil {
		return nil, err
	}

	return defaultBindings, nil
}
