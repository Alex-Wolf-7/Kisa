package game

import (
	"strconv"
	"strings"

	"github.com/Alex-Wolf-7/Kisa/plog"
)

type Phase int64

const (
	Phase_UNKNOWN = iota
	Phase_IN_PROGRESS
	Phase_NO_GAME
	Phase_CHAMP_SELECT
	Phase_LOBBY
	Phase_MATCHMAKING
)

func (p Phase) String() string {
	switch p {
	case Phase_UNKNOWN:
		return "Unknown"
	case Phase_IN_PROGRESS:
		return "In progress"
	case Phase_NO_GAME:
		return "No game"
	case Phase_CHAMP_SELECT:
		return "Champ select"
	case Phase_LOBBY:
		return "Lobby"
	case Phase_MATCHMAKING:
		return "Matchmaking"
	default:
		return "Phase not found"
	}
}

type Game struct {
	GameClient struct {
		ObserverServerIP   string `json:"observerServerIp"`
		ObserverServerPort int    `json:"observerServerPort"`
		Running            bool   `json:"running"`
		ServerIP           string `json:"serverIp"`
		ServerPort         int    `json:"serverPort"`
		Visible            bool   `json:"visible"`
	} `json:"gameClient"`
	GameData struct {
		GameID                   int64  `json:"gameId"`
		GameName                 string `json:"gameName"`
		IsCustomGame             bool   `json:"isCustomGame"`
		Password                 string `json:"password"`
		PlayerChampionSelections []struct {
			ChampionID           float64 `json:"championId"`
			SelectedSkinIndex    float64 `json:"selectedSkinIndex"`
			Spell1ID             float64 `json:"spell1Id"`
			Spell2ID             float64 `json:"spell2Id"`
			SummonerInternalName string  `json:"summonerInternalName"`
		} `json:"playerChampionSelections"`
		Queue struct {
			AllowablePremadeSizes   []interface{} `json:"allowablePremadeSizes"`
			AreFreeChampionsAllowed bool          `json:"areFreeChampionsAllowed"`
			AssetMutator            string        `json:"assetMutator"`
			Category                string        `json:"category"`
			ChampionsRequiredToPlay int           `json:"championsRequiredToPlay"`
			Description             string        `json:"description"`
			DetailedDescription     string        `json:"detailedDescription"`
			GameMode                string        `json:"gameMode"`
			GameTypeConfig          struct {
				AdvancedLearningQuests bool   `json:"advancedLearningQuests"`
				AllowTrades            bool   `json:"allowTrades"`
				BanMode                string `json:"banMode"`
				BanTimerDuration       int    `json:"banTimerDuration"`
				BattleBoost            bool   `json:"battleBoost"`
				CrossTeamChampionPool  bool   `json:"crossTeamChampionPool"`
				DeathMatch             bool   `json:"deathMatch"`
				DoNotRemove            bool   `json:"doNotRemove"`
				DuplicatePick          bool   `json:"duplicatePick"`
				ExclusivePick          bool   `json:"exclusivePick"`
				ID                     int    `json:"id"`
				LearningQuests         bool   `json:"learningQuests"`
				MainPickTimerDuration  int    `json:"mainPickTimerDuration"`
				MaxAllowableBans       int    `json:"maxAllowableBans"`
				Name                   string `json:"name"`
				OnboardCoopBeginner    bool   `json:"onboardCoopBeginner"`
				PickMode               string `json:"pickMode"`
				PostPickTimerDuration  int    `json:"postPickTimerDuration"`
				Reroll                 bool   `json:"reroll"`
				TeamChampionPool       bool   `json:"teamChampionPool"`
			} `json:"gameTypeConfig"`
			ID                                  int    `json:"id"`
			IsRanked                            bool   `json:"isRanked"`
			IsTeamBuilderManaged                bool   `json:"isTeamBuilderManaged"`
			IsTeamOnly                          bool   `json:"isTeamOnly"`
			LastToggledOffTime                  int    `json:"lastToggledOffTime"`
			LastToggledOnTime                   int    `json:"lastToggledOnTime"`
			MapID                               int    `json:"mapId"`
			MaxLevel                            int    `json:"maxLevel"`
			MaxSummonerLevelForFirstWinOfTheDay int    `json:"maxSummonerLevelForFirstWinOfTheDay"`
			MaximumParticipantListSize          int    `json:"maximumParticipantListSize"`
			MinLevel                            int    `json:"minLevel"`
			MinimumParticipantListSize          int    `json:"minimumParticipantListSize"`
			Name                                string `json:"name"`
			NumPlayersPerTeam                   int    `json:"numPlayersPerTeam"`
			QueueAvailability                   string `json:"queueAvailability"`
			QueueRewards                        struct {
				IsChampionPointsEnabled bool          `json:"isChampionPointsEnabled"`
				IsIPEnabled             bool          `json:"isIpEnabled"`
				IsXpEnabled             bool          `json:"isXpEnabled"`
				PartySizeIPRewards      []interface{} `json:"partySizeIpRewards"`
			} `json:"queueRewards"`
			RemovalFromGameAllowed      bool   `json:"removalFromGameAllowed"`
			RemovalFromGameDelayMinutes int    `json:"removalFromGameDelayMinutes"`
			ShortName                   string `json:"shortName"`
			ShowPositionSelector        bool   `json:"showPositionSelector"`
			SpectatorEnabled            bool   `json:"spectatorEnabled"`
			Type                        string `json:"type"`
		} `json:"queue"`
		SpectatorsAllowed bool `json:"spectatorsAllowed"`
		TeamOne           []struct {
			AccountID         float64 `json:"accountId"`
			AdjustmentFlags   float64 `json:"adjustmentFlags"`
			BotDifficulty     string  `json:"botDifficulty"`
			ClientInSynch     bool    `json:"clientInSynch"`
			GameCustomization struct {
				Regalia        string `json:"Regalia"`
				Perks          string `json:"perks"`
				SummonerEmotes string `json:"summonerEmotes"`
				WardSkin       string `json:"wardSkin"`
			} `json:"gameCustomization"`
			Index                   float64     `json:"index"`
			LastSelectedSkinIndex   float64     `json:"lastSelectedSkinIndex"`
			Locale                  interface{} `json:"locale"`
			Minor                   bool        `json:"minor"`
			OriginalAccountNumber   float64     `json:"originalAccountNumber"`
			OriginalPlatformID      string      `json:"originalPlatformId"`
			PartnerID               string      `json:"partnerId"`
			PickMode                float64     `json:"pickMode"`
			PickTurn                float64     `json:"pickTurn"`
			ProfileIconID           float64     `json:"profileIconId"`
			Puuid                   string      `json:"puuid"`
			QueueRating             float64     `json:"queueRating"`
			RankedTeamGuest         bool        `json:"rankedTeamGuest"`
			SelectedPosition        interface{} `json:"selectedPosition"`
			SelectedRole            interface{} `json:"selectedRole"`
			SummonerID              float64     `json:"summonerId"`
			SummonerInternalName    string      `json:"summonerInternalName"`
			SummonerName            string      `json:"summonerName"`
			TeamOwner               bool        `json:"teamOwner"`
			TeamParticipantID       interface{} `json:"teamParticipantId"`
			TeamRating              float64     `json:"teamRating"`
			TimeAddedToQueue        interface{} `json:"timeAddedToQueue"`
			TimeChampionSelectStart float64     `json:"timeChampionSelectStart"`
			TimeGameCreated         float64     `json:"timeGameCreated"`
			TimeMatchmakingStart    float64     `json:"timeMatchmakingStart"`
			VoterRating             float64     `json:"voterRating"`
		} `json:"teamOne"`
		TeamTwo []struct {
			BotDifficulty     string      `json:"botDifficulty"`
			BotSkillLevel     float64     `json:"botSkillLevel"`
			ChampionID        interface{} `json:"championId"`
			GameCustomization struct {
			} `json:"gameCustomization"`
			LastSelectedSkinIndex float64     `json:"lastSelectedSkinIndex"`
			Locale                interface{} `json:"locale"`
			PickMode              float64     `json:"pickMode"`
			PickTurn              float64     `json:"pickTurn"`
			Role                  interface{} `json:"role"`
			Spell1ID              interface{} `json:"spell1Id"`
			Spell2ID              interface{} `json:"spell2Id"`
			SummonerInternalName  string      `json:"summonerInternalName"`
			SummonerName          string      `json:"summonerName"`
			TeamID                string      `json:"teamId"`
		} `json:"teamTwo"`
	} `json:"gameData"`
	GameDodge struct {
		DodgeIds []interface{} `json:"dodgeIds"`
		Phase    string        `json:"phase"`
		State    string        `json:"state"`
	} `json:"gameDodge"`
	Map struct {
		Assets struct {
			ChampSelectBackgroundSound  string `json:"champ-select-background-sound"`
			ChampSelectFlyoutBackground string `json:"champ-select-flyout-background"`
			ChampSelectPlanningIntro    string `json:"champ-select-planning-intro"`
			GameSelectIconActive        string `json:"game-select-icon-active"`
			GameSelectIconActiveVideo   string `json:"game-select-icon-active-video"`
			GameSelectIconDefault       string `json:"game-select-icon-default"`
			GameSelectIconDisabled      string `json:"game-select-icon-disabled"`
			GameSelectIconHover         string `json:"game-select-icon-hover"`
			GameSelectIconIntroVideo    string `json:"game-select-icon-intro-video"`
			GameflowBackground          string `json:"gameflow-background"`
			GameselectButtonHoverSound  string `json:"gameselect-button-hover-sound"`
			IconDefeat                  string `json:"icon-defeat"`
			IconDefeatVideo             string `json:"icon-defeat-video"`
			IconEmpty                   string `json:"icon-empty"`
			IconHover                   string `json:"icon-hover"`
			IconLeaver                  string `json:"icon-leaver"`
			IconVictory                 string `json:"icon-victory"`
			IconVictoryVideo            string `json:"icon-victory-video"`
			MapNorth                    string `json:"map-north"`
			MapSouth                    string `json:"map-south"`
			MusicInqueueLoopSound       string `json:"music-inqueue-loop-sound"`
			PartiesBackground           string `json:"parties-background"`
			PostgameAmbienceLoopSound   string `json:"postgame-ambience-loop-sound"`
			ReadyCheckBackground        string `json:"ready-check-background"`
			ReadyCheckBackgroundSound   string `json:"ready-check-background-sound"`
			SfxAmbiencePregameLoopSound string `json:"sfx-ambience-pregame-loop-sound"`
			SocialIconLeaver            string `json:"social-icon-leaver"`
			SocialIconVictory           string `json:"social-icon-victory"`
		} `json:"assets"`
		CategorizedContentBundles struct {
		} `json:"categorizedContentBundles"`
		Description                         string `json:"description"`
		GameMode                            string `json:"gameMode"`
		GameModeName                        string `json:"gameModeName"`
		GameModeShortName                   string `json:"gameModeShortName"`
		GameMutator                         string `json:"gameMutator"`
		ID                                  int    `json:"id"`
		IsRGM                               bool   `json:"isRGM"`
		MapStringID                         string `json:"mapStringId"`
		Name                                string `json:"name"`
		PerPositionDisallowedSummonerSpells struct {
		} `json:"perPositionDisallowedSummonerSpells"`
		PerPositionRequiredSummonerSpells struct {
		} `json:"perPositionRequiredSummonerSpells"`
		PlatformID   string `json:"platformId"`
		PlatformName string `json:"platformName"`
		Properties   struct {
			SuppressRunesMasteriesPerks bool `json:"suppressRunesMasteriesPerks"`
		} `json:"properties"`
	} `json:"map"`
	Phase string `json:"phase"`
}

func (g *Game) GetPhase() Phase {
	switch g.Phase {
	case "":
		return Phase_NO_GAME
	case "InProgress":
		return Phase_IN_PROGRESS
	case "ChampSelect":
		return Phase_CHAMP_SELECT
	case "Lobby":
		return Phase_LOBBY
	case "Matchmaking":
		return Phase_MATCHMAKING
	default:
		return Phase_UNKNOWN
	}
}

func (g *Game) IsValid() bool {
	return g.GetPhase() != Phase_NO_GAME
}

func (g *Game) GetChampionNumberForSummoner(summonerInternalName string) string {
	plog.Debugf("Summoner: %s\n", summonerInternalName)
	plog.Debugf("PlayerChampionSelections: %v\n", g.GameData.PlayerChampionSelections)
	for _, selection := range g.GameData.PlayerChampionSelections {
		plog.Debugf("Is %s == %s?\n", summonerInternalName, selection.SummonerInternalName)
		if strings.ToLower(selection.SummonerInternalName) == strings.ToLower(summonerInternalName) {
			return strconv.FormatInt(int64(selection.ChampionID), 10)
		}
	}
	return ""
}
