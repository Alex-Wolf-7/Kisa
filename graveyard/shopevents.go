package gamesettings

type ShopEvents struct {
	EvtShopFocusSearch string `json:"evtShopFocusSearch,omitempty"`
	EvtShopSwitchTabs  string `json:"evtShopSwitchTabs,omitempty"`
}

func (seDefault *ShopEvents) ShopEventsDiff(seNew *ShopEvents) *ShopEvents {
	var evtShopFocusSearch, evtShopSwitchTabs string

	if seDefault.EvtShopFocusSearch != seNew.EvtShopFocusSearch {
		evtShopFocusSearch = seNew.EvtShopFocusSearch
	}

	if seDefault.EvtShopSwitchTabs != seNew.EvtShopSwitchTabs {
		evtShopSwitchTabs = seNew.EvtShopSwitchTabs
	}

	seDiff := &ShopEvents{
		EvtShopFocusSearch: evtShopFocusSearch,
		EvtShopSwitchTabs:  evtShopSwitchTabs,
	}

	empty := ShopEvents{}
	if *seDiff == empty {
		return nil
	} else {
		return seDiff
	}
}
