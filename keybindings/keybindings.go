package keybindings

import (
	"encoding/json"
)

const UNBOUND = "[\u003cUnbound\u003e]"

type KeyBindings struct {
	GameEvents GameEvents `json:"GameEvents"`
}

func (kb KeyBindings) MarshalJSON(omitEmpty bool) ([]byte, error) {
	if !omitEmpty {
		return json.Marshal(kb)
	}

	geMap, err := kb.GameEvents.ConvertToMap()

	var kbMap map[string]interface{}
	kbJSON, err := json.Marshal(kb)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(kbJSON, &kbMap)
	if err != nil {
		return nil, err
	}

	kbMap["GameEvents"] = geMap

	finalJSON, err := json.Marshal(kbMap)
	return finalJSON, err
}
