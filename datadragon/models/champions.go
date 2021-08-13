package models

import (
	"encoding/json"
	"fmt"
)

type ChampionListUnformatted struct {
	Type    string                 `json:"type"`
	Format  string                 `json:"format"`
	Version string                 `json:"version"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type Champion struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
		Hp                   int     `json:"hp"`
		Hpperlevel           int     `json:"hpperlevel"`
		Mp                   int     `json:"mp"`
		Mpperlevel           int     `json:"mpperlevel"`
		Movespeed            int     `json:"movespeed"`
		Armor                int     `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           int     `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          int     `json:"attackrange"`
		Hpregen              int     `json:"hpregen"`
		Hpregenperlevel      int     `json:"hpregenperlevel"`
		Mpregen              int     `json:"mpregen"`
		Mpregenperlevel      int     `json:"mpregenperlevel"`
		Crit                 int     `json:"crit"`
		Critperlevel         int     `json:"critperlevel"`
		Attackdamage         int     `json:"attackdamage"`
		Attackdamageperlevel int     `json:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed"`
	} `json:"stats"`
}

func BuildChampion(obj interface{}) (*Champion, error) {
	thing, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Unable to format champion into champion object")
	}

	champ := new(Champion)
	json.Unmarshal(thing, champ)

	return champ, nil
}
