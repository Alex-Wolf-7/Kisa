package gamesettings

type HUDEvents struct {
	EvtHoldShowScoreBoard       string `json:"evtHoldShowScoreBoard,omitempty"`
	EvtToggleDeathRecapShowcase string `json:"evtToggleDeathRecapShowcase,omitempty"`
	EvtToggleFPSAndLatency      string `json:"evtToggleFPSAndLatency,omitempty"`
	EvtToggleMouseClip          string `json:"evtToggleMouseClip,omitempty"`
	EvtTogglePlayerStats        string `json:"evtTogglePlayerStats,omitempty"`
}

func (heDefault *HUDEvents) HudEventsDiff(heNew *HUDEvents) *HUDEvents {
	var evtHoldShowScoreBoard, evtToggleDeathRecapShowcase, evtToggleFPSAndLatency, evtToggleMouseClip, evtTogglePlayerStats string

	if heDefault.EvtHoldShowScoreBoard != heNew.EvtHoldShowScoreBoard {
		evtHoldShowScoreBoard = heNew.EvtHoldShowScoreBoard
	}

	if heDefault.EvtToggleDeathRecapShowcase != heNew.EvtToggleDeathRecapShowcase {
		evtToggleDeathRecapShowcase = heNew.EvtToggleDeathRecapShowcase
	}

	if heDefault.EvtToggleFPSAndLatency != heNew.EvtToggleFPSAndLatency {
		evtToggleFPSAndLatency = heNew.EvtToggleFPSAndLatency
	}

	if heDefault.EvtToggleMouseClip != heNew.EvtToggleMouseClip {
		evtToggleMouseClip = heNew.EvtToggleMouseClip
	}

	if heDefault.EvtTogglePlayerStats != heNew.EvtTogglePlayerStats {
		evtTogglePlayerStats = heNew.EvtTogglePlayerStats
	}

	heDiff := &HUDEvents{
		EvtHoldShowScoreBoard:       evtHoldShowScoreBoard,
		EvtToggleDeathRecapShowcase: evtToggleDeathRecapShowcase,
		EvtToggleFPSAndLatency:      evtToggleFPSAndLatency,
		EvtToggleMouseClip:          evtToggleMouseClip,
		EvtTogglePlayerStats:        evtTogglePlayerStats,
	}

	empty := HUDEvents{}
	if *heDiff == empty {
		return nil
	} else {
		return heDiff
	}
}
