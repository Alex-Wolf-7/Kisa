package champions

import (
	"strings"

	"github.com/Alex-Wolf-7/Kisa/keybindings"
)

var champions = map[string]keybindings.KeyBindings{
	"default": {
		GameEvents: keybindings.GameEvents{
			SmartCastWithIndicator1: []string{"q"},
			SmartCastWithIndicator2: []string{"w"},
			SmartCastWithIndicator3: []string{"e"},
			SmartCastWithIndicator4: []string{"r"},
		},
	},
	"janna": {
		GameEvents: keybindings.GameEvents{
			SmartCast1: []string{"q"},

			SmartCastWithIndicator2: []string{"w"},

			NormalCast3: []string{"]"},
			SelfCast3:   []string{"Shift", "e"},
			SmartCast3:  []string{"e"},

			SmartCast4: []string{"r"},
		},
	},
}

func GetChampion(champion string) keybindings.KeyBindings {
	if keyBindings, ok := champions[strings.ToLower(champion)]; ok {
		return keyBindings
	} else {
		return champions["default"]
	}
}
