package keybindings

const UNBOUND = "[\u003cUnbound\u003e]"

type KeyBindings struct {
	Champion   string     `json:"-"`
	GameEvents GameEvents `json:"GameEvents"`
}
