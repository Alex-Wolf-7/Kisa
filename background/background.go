package background

import (
	"fmt"
	"time"

	"github.com/Alex-Wolf-7/Kisa/constants"
	"github.com/Alex-Wolf-7/Kisa/datadragon/models"
	"github.com/Alex-Wolf-7/Kisa/game"
	"github.com/Alex-Wolf-7/Kisa/lolclient"
	"github.com/Alex-Wolf-7/Kisa/plog"
	"github.com/Alex-Wolf-7/Kisa/settingsdb"
	"github.com/Alex-Wolf-7/Kisa/summoner"
)

type Background struct {
	lolClient   *lolclient.LoLClient
	settingsDB  *settingsdb.SettingsDB
	summoner    *summoner.Summoner
	championMap map[string]*models.Champion
	stop        bool
}

func NewBackground(
	lolClient *lolclient.LoLClient,
	settingsDB *settingsdb.SettingsDB,
	summoner *summoner.Summoner,
	championMap map[string]*models.Champion,
) *Background {

	return &Background{
		lolClient:   lolClient,
		settingsDB:  settingsDB,
		summoner:    summoner,
		championMap: championMap,
		stop:        false,
	}
}

func (b *Background) Loop() error {
	for !b.stop {
		champion, err := b.WaitForGameStart()
		if err != nil {
			return fmt.Errorf("failed to wait for game start: %s", err.Error())
		}

		gameSettings, err := b.settingsDB.GetSettings(champion.Name)
		if err != nil {
			return fmt.Errorf("unable to get settings for chosen champion: %s", err.Error())
		}
		if gameSettings == nil {
			gameSettings, err = b.settingsDB.GetDefaultSettings()
			if err != nil {
				return fmt.Errorf("unable to get settings for chosen champion: %s", err.Error())
			}
			plog.Infof("No champion-specific settings: using default settings\n")
		} else {
			plog.Infof("Champion settings obtained\n")
		}

		err = b.lolClient.PatchGameSettingsMultiple(gameSettings)
		if err != nil {
			return fmt.Errorf("unable to change game settings: %s", err.Error())
		}

		err = b.WaitForGameEnd()
		if err != nil {
			return fmt.Errorf("failed to wait for game to end: %s", err.Error())
		}
	}

	return nil
}

func (b *Background) WaitForGameStart() (*models.Champion, error) {
	for {
		currentGame, err := b.lolClient.GetGameSession()
		if err != nil {
			return nil, err
		}

		if currentGame.GetPhase() == game.Phase_IN_PROGRESS {
			champNum := currentGame.GetChampionNumberForSummoner(b.summoner.InternalName)
			for champNum == "" {
				plog.Debugf("Failed to get current champion: trying again in 1 second")
				time.Sleep(1 * time.Second)
				champNum = currentGame.GetChampionNumberForSummoner(b.summoner.InternalName)
			}
			champ, ok := b.championMap[champNum]
			if !ok {
				return nil, fmt.Errorf("unable to get champion from champion number %s", champNum)
			}
			plog.Infof("Game started. Chosen champ: %s\n", champ.Name)
			return champ, nil
		} else {
			plog.Periodicf("No game started\n")
			time.Sleep(constants.CHECK_IF_GAME_STARTED_TIME)
		}
	}
}

func (b *Background) WaitForGameEnd() error {
	for {
		currentGame, err := b.lolClient.GetGameSession()
		if err != nil {
			return err
		}

		if currentGame.GetPhase() == game.Phase_IN_PROGRESS {
			time.Sleep(constants.CHECK_IF_GAME_OVER_TIME)
		} else {
			plog.Infof("Game finished\n")
			return nil
		}
	}
}

func (b *Background) Stop() {
	b.stop = true
}
