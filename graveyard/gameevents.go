package gamesettings

import (
	"encoding/json"
	"fmt"
)

type GameEvents struct {
	EvntPlayerPing                                string `json:"evntPlayerPing,omitempty"`
	EvntPlayerPingCursor                          string `json:"evntPlayerPingCursor,omitempty"`
	EvntPlayerPingCursorDanger                    string `json:"evntPlayerPingCursorDanger,omitempty"`
	EvntPlayerPingDanger                          string `json:"evntPlayerPingDanger,omitempty"`
	EvtCameraLockToggle                           string `json:"evtCameraLockToggle,omitempty"`
	EvtCameraSnap                                 string `json:"evtCameraSnap,omitempty"`
	EvtCastAvatarSpell1                           string `json:"evtCastAvatarSpell1,omitempty"`
	EvtCastAvatarSpell2                           string `json:"evtCastAvatarSpell2,omitempty"`
	EvtCastSpell1                                 string `json:"evtCastSpell1,omitempty"`
	EvtCastSpell2                                 string `json:"evtCastSpell2,omitempty"`
	EvtCastSpell3                                 string `json:"evtCastSpell3,omitempty"`
	EvtCastSpell4                                 string `json:"evtCastSpell4,omitempty"`
	EvtChampMasteryDisplay                        string `json:"evtChampMasteryDisplay,omitempty"`
	EvtChampionOnly                               string `json:"evtChampionOnly,omitempty"`
	EvtChatHistory                                string `json:"evtChatHistory,omitempty"`
	EvtDragScrollLock                             string `json:"evtDragScrollLock,omitempty"`
	EvtDrawHud                                    string `json:"evtDrawHud,omitempty"`
	EvtEmoteDance                                 string `json:"evtEmoteDance,omitempty"`
	EvtEmoteJoke                                  string `json:"evtEmoteJoke,omitempty"`
	EvtEmoteLaugh                                 string `json:"evtEmoteLaugh,omitempty"`
	EvtEmoteTaunt                                 string `json:"evtEmoteTaunt,omitempty"`
	EvtEmoteToggle                                string `json:"evtEmoteToggle,omitempty"`
	EvtLevelSpell1                                string `json:"evtLevelSpell1,omitempty"`
	EvtLevelSpell2                                string `json:"evtLevelSpell2,omitempty"`
	EvtLevelSpell3                                string `json:"evtLevelSpell3,omitempty"`
	EvtLevelSpell4                                string `json:"evtLevelSpell4,omitempty"`
	EvtNormalCastAvatarSpell1                     string `json:"evtNormalCastAvatarSpell1,omitempty"`
	EvtNormalCastAvatarSpell2                     string `json:"evtNormalCastAvatarSpell2,omitempty"`
	EvtNormalCastItem1                            string `json:"evtNormalCastItem1,omitempty"`
	EvtNormalCastItem2                            string `json:"evtNormalCastItem2,omitempty"`
	EvtNormalCastItem3                            string `json:"evtNormalCastItem3,omitempty"`
	EvtNormalCastItem4                            string `json:"evtNormalCastItem4,omitempty"`
	EvtNormalCastItem5                            string `json:"evtNormalCastItem5,omitempty"`
	EvtNormalCastItem6                            string `json:"evtNormalCastItem6,omitempty"`
	EvtNormalCastSpell1                           string `json:"evtNormalCastSpell1,omitempty"`
	EvtNormalCastSpell2                           string `json:"evtNormalCastSpell2,omitempty"`
	EvtNormalCastSpell3                           string `json:"evtNormalCastSpell3,omitempty"`
	EvtNormalCastSpell4                           string `json:"evtNormalCastSpell4,omitempty"`
	EvtNormalCastVisionItem                       string `json:"evtNormalCastVisionItem,omitempty"`
	EvtOnUIMouse4Pan                              string `json:"evtOnUIMouse4Pan,omitempty"`
	EvtOpenShop                                   string `json:"evtOpenShop,omitempty"`
	EvtPetMoveClick                               string `json:"evtPetMoveClick,omitempty"`
	EvtPlayerAttackMove                           string `json:"evtPlayerAttackMove,omitempty"`
	EvtPlayerAttackMoveClick                      string `json:"evtPlayerAttackMoveClick,omitempty"`
	EvtPlayerAttackOnlyClick                      string `json:"evtPlayerAttackOnlyClick,omitempty"`
	EvtPlayerHoldPosition                         string `json:"evtPlayerHoldPosition,omitempty"`
	EvtPlayerMoveClick                            string `json:"evtPlayerMoveClick,omitempty"`
	EvtPlayerPingAreaIsWarded                     string `json:"evtPlayerPingAreaIsWarded,omitempty"`
	EvtPlayerPingComeHere                         string `json:"evtPlayerPingComeHere,omitempty"`
	EvtPlayerPingMIA                              string `json:"evtPlayerPingMIA,omitempty"`
	EvtPlayerPingOMW                              string `json:"evtPlayerPingOMW,omitempty"`
	EvtPlayerPingRadialDanger                     string `json:"evtPlayerPingRadialDanger,omitempty"`
	EvtPlayerStopPosition                         string `json:"evtPlayerStopPosition,omitempty"`
	EvtPushToTalk                                 string `json:"evtPushToTalk,omitempty"`
	EvtRadialEmoteInstantOpen                     string `json:"evtRadialEmoteInstantOpen,omitempty"`
	EvtRadialEmoteOpen                            string `json:"evtRadialEmoteOpen,omitempty"`
	EvtRadialEmotePlaySlot0                       string `json:"evtRadialEmotePlaySlot0,omitempty"`
	EvtRadialEmotePlaySlot1                       string `json:"evtRadialEmotePlaySlot1,omitempty"`
	EvtRadialEmotePlaySlot2                       string `json:"evtRadialEmotePlaySlot2,omitempty"`
	EvtRadialEmotePlaySlot3                       string `json:"evtRadialEmotePlaySlot3,omitempty"`
	EvtRadialEmotePlaySlot4                       string `json:"evtRadialEmotePlaySlot4,omitempty"`
	EvtScrollDown                                 string `json:"evtScrollDown,omitempty"`
	EvtScrollLeft                                 string `json:"evtScrollLeft,omitempty"`
	EvtScrollRight                                string `json:"evtScrollRight,omitempty"`
	EvtScrollUp                                   string `json:"evtScrollUp,omitempty"`
	EvtSelectAlly1                                string `json:"evtSelectAlly1,omitempty"`
	EvtSelectAlly2                                string `json:"evtSelectAlly2,omitempty"`
	EvtSelectAlly3                                string `json:"evtSelectAlly3,omitempty"`
	EvtSelectAlly4                                string `json:"evtSelectAlly4,omitempty"`
	EvtSelectSelf                                 string `json:"evtSelectSelf,omitempty"`
	EvtSelfCastAvatarSpell1                       string `json:"evtSelfCastAvatarSpell1,omitempty"`
	EvtSelfCastAvatarSpell2                       string `json:"evtSelfCastAvatarSpell2,omitempty"`
	EvtSelfCastItem1                              string `json:"evtSelfCastItem1,omitempty"`
	EvtSelfCastItem2                              string `json:"evtSelfCastItem2,omitempty"`
	EvtSelfCastItem3                              string `json:"evtSelfCastItem3,omitempty"`
	EvtSelfCastItem4                              string `json:"evtSelfCastItem4,omitempty"`
	EvtSelfCastItem5                              string `json:"evtSelfCastItem5,omitempty"`
	EvtSelfCastItem6                              string `json:"evtSelfCastItem6,omitempty"`
	EvtSelfCastSpell1                             string `json:"evtSelfCastSpell1,omitempty"`
	EvtSelfCastSpell2                             string `json:"evtSelfCastSpell2,omitempty"`
	EvtSelfCastSpell3                             string `json:"evtSelfCastSpell3,omitempty"`
	EvtSelfCastSpell4                             string `json:"evtSelfCastSpell4,omitempty"`
	EvtSelfCastVisionItem                         string `json:"evtSelfCastVisionItem,omitempty"`
	EvtShowCharacterMenu                          string `json:"evtShowCharacterMenu,omitempty"`
	EvtShowHealthBars                             string `json:"evtShowHealthBars,omitempty"`
	EvtShowScoreBoard                             string `json:"evtShowScoreBoard,omitempty"`
	EvtShowSummonerNames                          string `json:"evtShowSummonerNames,omitempty"`
	EvtShowVoicePanel                             string `json:"evtShowVoicePanel,omitempty"`
	EvtSmartCastAvatarSpell1                      string `json:"evtSmartCastAvatarSpell1,omitempty"`
	EvtSmartCastAvatarSpell2                      string `json:"evtSmartCastAvatarSpell2,omitempty"`
	EvtSmartCastItem1                             string `json:"evtSmartCastItem1,omitempty"`
	EvtSmartCastItem2                             string `json:"evtSmartCastItem2,omitempty"`
	EvtSmartCastItem3                             string `json:"evtSmartCastItem3,omitempty"`
	EvtSmartCastItem4                             string `json:"evtSmartCastItem4,omitempty"`
	EvtSmartCastItem5                             string `json:"evtSmartCastItem5,omitempty"`
	EvtSmartCastItem6                             string `json:"evtSmartCastItem6,omitempty"`
	EvtSmartCastSpell1                            string `json:"evtSmartCastSpell1,omitempty"`
	EvtSmartCastSpell2                            string `json:"evtSmartCastSpell2,omitempty"`
	EvtSmartCastSpell3                            string `json:"evtSmartCastSpell3,omitempty"`
	EvtSmartCastSpell4                            string `json:"evtSmartCastSpell4,omitempty"`
	EvtSmartCastVisionItem                        string `json:"evtSmartCastVisionItem,omitempty"`
	EvtSmartCastWithIndicatorAvatarSpell1         string `json:"evtSmartCastWithIndicatorAvatarSpell1,omitempty"`
	EvtSmartCastWithIndicatorAvatarSpell2         string `json:"evtSmartCastWithIndicatorAvatarSpell2,omitempty"`
	EvtSmartCastWithIndicatorItem1                string `json:"evtSmartCastWithIndicatorItem1,omitempty"`
	EvtSmartCastWithIndicatorItem2                string `json:"evtSmartCastWithIndicatorItem2,omitempty"`
	EvtSmartCastWithIndicatorItem3                string `json:"evtSmartCastWithIndicatorItem3,omitempty"`
	EvtSmartCastWithIndicatorItem4                string `json:"evtSmartCastWithIndicatorItem4,omitempty"`
	EvtSmartCastWithIndicatorItem5                string `json:"evtSmartCastWithIndicatorItem5,omitempty"`
	EvtSmartCastWithIndicatorItem6                string `json:"evtSmartCastWithIndicatorItem6,omitempty"`
	EvtSmartCastWithIndicatorSpell1               string `json:"evtSmartCastWithIndicatorSpell1,omitempty"`
	EvtSmartCastWithIndicatorSpell2               string `json:"evtSmartCastWithIndicatorSpell2,omitempty"`
	EvtSmartCastWithIndicatorSpell3               string `json:"evtSmartCastWithIndicatorSpell3,omitempty"`
	EvtSmartCastWithIndicatorSpell4               string `json:"evtSmartCastWithIndicatorSpell4,omitempty"`
	EvtSmartCastWithIndicatorVisionItem           string `json:"evtSmartCastWithIndicatorVisionItem,omitempty"`
	EvtSmartPlusSelfCastAvatarSpell1              string `json:"evtSmartPlusSelfCastAvatarSpell1,omitempty"`
	EvtSmartPlusSelfCastAvatarSpell2              string `json:"evtSmartPlusSelfCastAvatarSpell2,omitempty"`
	EvtSmartPlusSelfCastItem1                     string `json:"evtSmartPlusSelfCastItem1,omitempty"`
	EvtSmartPlusSelfCastItem2                     string `json:"evtSmartPlusSelfCastItem2,omitempty"`
	EvtSmartPlusSelfCastItem3                     string `json:"evtSmartPlusSelfCastItem3,omitempty"`
	EvtSmartPlusSelfCastItem4                     string `json:"evtSmartPlusSelfCastItem4,omitempty"`
	EvtSmartPlusSelfCastItem5                     string `json:"evtSmartPlusSelfCastItem5,omitempty"`
	EvtSmartPlusSelfCastItem6                     string `json:"evtSmartPlusSelfCastItem6,omitempty"`
	EvtSmartPlusSelfCastSpell1                    string `json:"evtSmartPlusSelfCastSpell1,omitempty"`
	EvtSmartPlusSelfCastSpell2                    string `json:"evtSmartPlusSelfCastSpell2,omitempty"`
	EvtSmartPlusSelfCastSpell3                    string `json:"evtSmartPlusSelfCastSpell3,omitempty"`
	EvtSmartPlusSelfCastSpell4                    string `json:"evtSmartPlusSelfCastSpell4,omitempty"`
	EvtSmartPlusSelfCastVisionItem                string `json:"evtSmartPlusSelfCastVisionItem,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorAvatarSpell1 string `json:"evtSmartPlusSelfCastWithIndicatorAvatarSpell1,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorAvatarSpell2 string `json:"evtSmartPlusSelfCastWithIndicatorAvatarSpell2,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem1        string `json:"evtSmartPlusSelfCastWithIndicatorItem1,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem2        string `json:"evtSmartPlusSelfCastWithIndicatorItem2,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem3        string `json:"evtSmartPlusSelfCastWithIndicatorItem3,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem4        string `json:"evtSmartPlusSelfCastWithIndicatorItem4,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem5        string `json:"evtSmartPlusSelfCastWithIndicatorItem5,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorItem6        string `json:"evtSmartPlusSelfCastWithIndicatorItem6,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorSpell1       string `json:"evtSmartPlusSelfCastWithIndicatorSpell1,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorSpell2       string `json:"evtSmartPlusSelfCastWithIndicatorSpell2,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorSpell3       string `json:"evtSmartPlusSelfCastWithIndicatorSpell3,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorSpell4       string `json:"evtSmartPlusSelfCastWithIndicatorSpell4,omitempty"`
	EvtSmartPlusSelfCastWithIndicatorVisionItem   string `json:"evtSmartPlusSelfCastWithIndicatorVisionItem,omitempty"`
	EvtSysMenu                                    string `json:"evtSysMenu,omitempty"`
	EvtToggleMinionHealthBars                     string `json:"evtToggleMinionHealthBars,omitempty"`
	EvtUseItem1                                   string `json:"evtUseItem1,omitempty"`
	EvtUseItem2                                   string `json:"evtUseItem2,omitempty"`
	EvtUseItem3                                   string `json:"evtUseItem3,omitempty"`
	EvtUseItem4                                   string `json:"evtUseItem4,omitempty"`
	EvtUseItem5                                   string `json:"evtUseItem5,omitempty"`
	EvtUseItem6                                   string `json:"evtUseItem6,omitempty"`
	EvtUseItem7                                   string `json:"evtUseItem7,omitempty"`
	EvtUseVisionItem                              string `json:"evtUseVisionItem,omitempty"`
}

func (ge1 *GameEvents) GameEventsDiff(ge2 *GameEvents) (*GameEvents, error) {
	ge1Byte, err := json.Marshal(ge1)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal GameEvent 1 into byte array: %s", err)
	}

	ge2Byte, err := json.Marshal(ge2)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal GameEvent 2 into byte array: %s", err)
	}

	ge1Map := make(map[string]string)
	err = json.Unmarshal(ge1Byte, &ge1Map)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal GameEvent 1 into map: %s", err)
	}

	ge2Map := make(map[string]string)
	err = json.Unmarshal(ge2Byte, &ge2Map)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal GameEvent 2 into map: %s", err)
	}

	for key, value := range ge1Map {
		if ge2Map[key] == value {
			delete(ge2Map, key)
		}
	}

	geByte, err := json.Marshal(ge2Map)
	if err != nil {
		return nil, fmt.Errorf("Unable to re-marshal resultant Game Event map into byte array: %s", err)
	}

	geFinal := new(GameEvents)
	err = json.Unmarshal(geByte, &geFinal)
	if err != nil {
		return nil, fmt.Errorf("Unable to re-unmarshal resultant GameEvent into struct: %s", err)
	}

	empty := GameEvents{}
	if *geFinal == empty {
		return nil, nil
	} else {
		return geFinal, nil
	}
}
