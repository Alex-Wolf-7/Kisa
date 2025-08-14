package keybindings

const UNBOUND = "[\u003cUnbound\u003e]"

type KeyBindings struct {
	GameEvents GameEvents `json:"GameEvents"`
}
