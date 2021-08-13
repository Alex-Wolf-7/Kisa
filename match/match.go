package match

type Phase int64

const (
	Phase_UNKNOWN = iota
	Phase_BAN_PICK
	Phase_FINALIZATION
)

type Match struct {
	Actions [][]struct {
		ActorCellID  int64  `json:"actorCellId"`
		ChampionID   int64  `json:"championId"`
		Completed    bool   `json:"completed"`
		ID           int64  `json:"id"`
		IsAllyAction bool   `json:"isAllyAction"`
		IsInProgress bool   `json:"isInProgress"`
		PickTurn     int64  `json:"pickTurn"`
		Type         string `json:"type"`
	} `json:"actions"`
	AllowBattleBoost    bool `json:"allowBattleBoost"`
	AllowDuplicatePicks bool `json:"allowDuplicatePicks"`
	AllowLockedEvents   bool `json:"allowLockedEvents"`
	AllowRerolling      bool `json:"allowRerolling"`
	AllowSkinSelection  bool `json:"allowSkinSelection"`
	Bans                struct {
		MyTeamBans    []interface{} `json:"myTeamBans"`
		NumBans       int64         `json:"numBans"`
		TheirTeamBans []interface{} `json:"theirTeamBans"`
	} `json:"bans"`
	BenchChampionIds   []interface{} `json:"benchChampionIds"`
	BenchEnabled       bool          `json:"benchEnabled"`
	BoostableSkinCount int64         `json:"boostableSkinCount"`
	ChatDetails        struct {
		ChatRoomName     string `json:"chatRoomName"`
		ChatRoomPassword string `json:"chatRoomPassword"`
	} `json:"chatDetails"`
	Counter              int64 `json:"counter"`
	EntitledFeatureState struct {
		AdditionalRerolls int64         `json:"additionalRerolls"`
		UnlockedSkinIds   []interface{} `json:"unlockedSkinIds"`
	} `json:"entitledFeatureState"`
	GameID               int64  `json:"gameId"`
	HasSimultaneousBans  bool   `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool   `json:"hasSimultaneousPicks"`
	IsCustomGame         bool   `json:"isCustomGame"`
	IsSpectating         bool   `json:"isSpectating"`
	LocalPlayerCellID    int64  `json:"localPlayerCellId"`
	LockedEventIndex     int64  `json:"lockedEventIndex"`
	Message              string `json:"message"`
	MyTeam               []struct {
		AssignedPosition    string `json:"assignedPosition"`
		CellID              int64  `json:"cellId"`
		ChampionID          int64  `json:"championId"`
		ChampionPickIntent  int64  `json:"championPickIntent"`
		EntitledFeatureType string `json:"entitledFeatureType"`
		SelectedSkinID      int64  `json:"selectedSkinId"`
		Spell1ID            int64  `json:"spell1Id"`
		Spell2ID            int64  `json:"spell2Id"`
		SummonerID          int64  `json:"summonerId"`
		Team                int64  `json:"team"`
		WardSkinID          int64  `json:"wardSkinId"`
	} `json:"myTeam"`
	RecoveryCounter    int64 `json:"recoveryCounter"`
	RerollsRemaining   int64 `json:"rerollsRemaining"`
	SkipChampionSelect bool  `json:"skipChampionSelect"`
	TheirTeam          []struct {
		AssignedPosition    string `json:"assignedPosition"`
		CellID              int64  `json:"cellId"`
		ChampionID          int64  `json:"championId"`
		ChampionPickIntent  int64  `json:"championPickIntent"`
		EntitledFeatureType string `json:"entitledFeatureType"`
		SelectedSkinID      int64  `json:"selectedSkinId"`
		Spell1ID            uint64 `json:"spell1Id"`
		Spell2ID            uint64 `json:"spell2Id"`
		SummonerID          int64  `json:"summonerId"`
		Team                int64  `json:"team"`
		WardSkinID          int64  `json:"wardSkinId"`
	} `json:"theirTeam"`
	Timer struct {
		AdjustedTimeLeftInPhase int64  `json:"adjustedTimeLeftInPhase"`
		InternalNowInEpochMs    int64  `json:"internalNowInEpochMs"`
		IsInfinite              bool   `json:"isInfinite"`
		Phase                   string `json:"phase"`
		TotalTimeInPhase        int64  `json:"totalTimeInPhase"`
	} `json:"timer"`
	Trades []interface{} `json:"trades"`
}

func (m *Match) GetPhase() Phase {
	switch m.Timer.Phase {
	case "BAN_PICK":
		return Phase_BAN_PICK
	case "FINALIZATION":
		return Phase_FINALIZATION
	default:
		return Phase_UNKNOWN
	}
}

func (m *Match) IsValid() bool {
	return m.Message != "No active delegate"
}
