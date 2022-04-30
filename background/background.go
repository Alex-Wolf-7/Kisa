package background

import (
	"strconv"
	"time"

	"github.com/Alex-Wolf-7/Kisa/champions"
	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/datadragon/models"
	"github.com/Alex-Wolf-7/Kisa/game"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/plog"
)

type Background struct {
	lolClient   *lolclient.LoLClient
	championMap map[string]*models.Champion
}

func NewBackground(
	lolClient *lolclient.LoLClient,
	championMap map[string]*models.Champion,
) *Background {

	return &Background{
		lolClient:   lolClient,
		championMap: championMap,
	}
}

func (b *Background) Loop() error {
	var champNum int
	for {
		currentGame, err := b.lolClient.GetGameSession()
		if err != nil {
			return err
		}

		switch currentGame.GetPhase() {
		case game.Phase_CHAMP_SELECT:
			// In champ select: keep getting most up-to-date locked champion and set keybindings
			champNum, err = b.lolClient.GetLockedChampion()
			if err != nil {
				return err
			}

			if champNum != 0 {
				champ := b.championMap[strconv.Itoa(champNum)]
				err = b.lolClient.PatchKeyBindings(champions.GetChampion(champ.Name))
				if err != nil {
					plog.ErrorfWithBackup("unable to set keybindings", "unable to set keybindings to champion %s", champ.Name)
				}
			}
			time.Sleep(constants.CHECK_IF_GAME_STARTED_TIME)
			continue
		case game.Phase_IN_PROGRESS:
			// Game started: be done
			return nil
		default:
			// No champ select or game started: wait for game start
			time.Sleep(constants.CHECK_IF_CHAMP_SELECT_TIME)
		}
	}
}
