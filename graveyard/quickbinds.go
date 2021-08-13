package gamesettings

import "fmt"

type Quickbinds struct {
	EvtCastAvatarSpell1Smart *bool `json:"evtCastAvatarSpell1smart,omitempty"`
	EvtCastAvatarSpell2Smart *bool `json:"evtCastAvatarSpell2smart,omitempty"`
	EvtCastSpell1Smart       *bool `json:"evtCastSpell1smart,omitempty"`
	EvtCastSpell2Smart       *bool `json:"evtCastSpell2smart,omitempty"`
	EvtCastSpell3Smart       *bool `json:"evtCastSpell3smart,omitempty"`
	EvtCastSpell4Smart       *bool `json:"evtCastSpell4smart,omitempty"`
	EvtUseItem1Smart         *bool `json:"evtUseItem1smart,omitempty"`
	EvtUseItem2Smart         *bool `json:"evtUseItem2smart,omitempty"`
	EvtUseItem3Smart         *bool `json:"evtUseItem3smart,omitempty"`
	EvtUseItem4Smart         *bool `json:"evtUseItem4smart,omitempty"`
	EvtUseItem5Smart         *bool `json:"evtUseItem5smart,omitempty"`
	EvtUseItem6Smart         *bool `json:"evtUseItem6smart,omitempty"`
	EvtUseVisionItemsmart    *bool `json:"evtUseVisionItemsmart,omitempty"`
}

func (qbDefault *Quickbinds) QuickbindsDiff(qbNew *Quickbinds) *Quickbinds {
	var evtCastAvatarSpell1Smart, evtCastAvatarSpell2Smart, evtCastSpell1Smart, evtCastSpell2Smart, evtCastSpell3Smart, evtCastSpell4Smart, evtUseItem1Smart, evtUseItem2Smart, evtUseItem3Smart, evtUseItem4Smart, evtUseItem5Smart, evtUseItem6Smart, evtUseVisionItemsmart *bool
	fmt.Println("Default:", *qbDefault)
	fmt.Println("Jannaaa:", *qbNew)

	if *qbDefault.EvtCastAvatarSpell1Smart != *qbNew.EvtCastAvatarSpell1Smart {
		evtCastAvatarSpell1Smart = qbNew.EvtCastAvatarSpell1Smart
	}

	if *qbDefault.EvtCastAvatarSpell2Smart != *qbNew.EvtCastAvatarSpell2Smart {
		evtCastAvatarSpell2Smart = qbNew.EvtCastAvatarSpell2Smart
	}

	if *qbDefault.EvtCastSpell1Smart != *qbNew.EvtCastSpell1Smart {
		evtCastSpell1Smart = qbNew.EvtCastSpell1Smart
	}

	if *qbDefault.EvtCastSpell2Smart != *qbNew.EvtCastSpell2Smart {
		evtCastSpell2Smart = qbNew.EvtCastSpell2Smart
	}

	if *qbDefault.EvtCastSpell3Smart != *qbNew.EvtCastSpell3Smart {
		evtCastSpell3Smart = qbNew.EvtCastSpell3Smart
	}

	if *qbDefault.EvtCastSpell4Smart != *qbNew.EvtCastSpell4Smart {
		evtCastSpell4Smart = qbNew.EvtCastSpell4Smart
	}

	if *qbDefault.EvtUseItem1Smart != *qbNew.EvtUseItem1Smart {
		evtUseItem1Smart = qbNew.EvtUseItem1Smart
	}

	if *qbDefault.EvtUseItem2Smart != *qbNew.EvtUseItem2Smart {
		evtUseItem2Smart = qbNew.EvtUseItem2Smart
	}

	if *qbDefault.EvtUseItem3Smart != *qbNew.EvtUseItem3Smart {
		evtUseItem3Smart = qbNew.EvtUseItem3Smart
	}

	if *qbDefault.EvtUseItem4Smart != *qbNew.EvtUseItem4Smart {
		evtUseItem4Smart = qbNew.EvtUseItem4Smart
	}

	if *qbDefault.EvtUseItem5Smart != *qbNew.EvtUseItem5Smart {
		evtUseItem5Smart = qbNew.EvtUseItem5Smart
	}

	if *qbDefault.EvtUseItem6Smart != *qbNew.EvtUseItem6Smart {
		evtUseItem6Smart = qbNew.EvtUseItem6Smart
	}

	if *qbDefault.EvtUseVisionItemsmart != *qbNew.EvtUseVisionItemsmart {
		evtUseVisionItemsmart = qbNew.EvtUseVisionItemsmart
	}

	qbDiff := &Quickbinds{
		EvtCastAvatarSpell1Smart: evtCastAvatarSpell1Smart,
		EvtCastAvatarSpell2Smart: evtCastAvatarSpell2Smart,
		EvtCastSpell1Smart:       evtCastSpell1Smart,
		EvtCastSpell2Smart:       evtCastSpell2Smart,
		EvtCastSpell3Smart:       evtCastSpell3Smart,
		EvtCastSpell4Smart:       evtCastSpell4Smart,
		EvtUseItem1Smart:         evtUseItem1Smart,
		EvtUseItem2Smart:         evtUseItem2Smart,
		EvtUseItem3Smart:         evtUseItem3Smart,
		EvtUseItem4Smart:         evtUseItem4Smart,
		EvtUseItem5Smart:         evtUseItem5Smart,
		EvtUseItem6Smart:         evtUseItem6Smart,
		EvtUseVisionItemsmart:    evtUseVisionItemsmart,
	}

	empty := Quickbinds{}
	if *qbDiff == empty {
		return nil
	} else {
		return qbDiff
	}
}
