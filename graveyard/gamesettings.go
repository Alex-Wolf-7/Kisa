package gamesettings

import (
	"github.com/Alex-Wolf-7/Kisa/plog"
)

type GameSettings struct {
	GameEvents *GameEvents `json:"GameEvents,omitempty"`
	HUDEvents  *HUDEvents  `json:"HUDEvents,omitempty"`
	Quickbinds *Quickbinds `json:"Quickbinds,omitempty"`
	ShopEvents *ShopEvents `json:"ShopEvents,omitempty"`
}

func (gsDefault *GameSettings) GameSettingsDiff(gsNew *GameSettings) *GameSettings {
	gameEvents, err := gsDefault.GameEvents.GameEventsDiff(gsNew.GameEvents)
	if err != nil {
		plog.ErrorfWithBackup("Unable to compare new game settings to old", "Unable to get the difference in Game Events: %s", err)
	}

	hudEvents := gsDefault.HUDEvents.HudEventsDiff(gsNew.HUDEvents)
	quickbinds := gsDefault.Quickbinds.QuickbindsDiff(gsNew.Quickbinds)
	shopEvents := gsDefault.ShopEvents.ShopEventsDiff(gsNew.ShopEvents)

	if gameEvents == nil && hudEvents == nil && quickbinds == nil && shopEvents == nil {
		return nil
	} else {
		return &GameSettings{
			GameEvents: gameEvents,
			HUDEvents:  hudEvents,
			Quickbinds: quickbinds,
			ShopEvents: shopEvents,
		}
	}
}
