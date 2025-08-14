package gamesettings

type GameSettings map[string]map[string]interface{}

func NewGameSettings() GameSettings {
	return make(GameSettings)
}

func (old GameSettings) GetChanges(new GameSettings) map[string]map[string]bool {
	addMap := make(map[string]map[string]bool)
	for outerKey, oldInnerMap := range old {
		for innerKey, oldValue := range oldInnerMap {
			newValue, ok := new[outerKey][innerKey]
			if !ok || newValue != oldValue {
				addMap[outerKey] = safeMapBoolAdd(addMap[outerKey], innerKey, true)
			}
		}
	}

	for outerKey, newInnerMap := range new {
		for innerKey, newValue := range newInnerMap {
			oldValue, ok := old[outerKey][innerKey]
			if !ok || newValue != oldValue {
				addMap[outerKey] = safeMapBoolAdd(addMap[outerKey], innerKey, true)
			}
		}
	}

	return addMap
}
