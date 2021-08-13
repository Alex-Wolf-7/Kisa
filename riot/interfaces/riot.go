package interfaces

import (
	"github.com/Alex-Wolf-7/Kisa/match"
	"github.com/Alex-Wolf-7/Kisa/riot/models"
	"github.com/Alex-Wolf-7/Kisa/summoner"
)

type RiotClient interface {
	GetSummonerByName(name string) (*summoner.Summoner, error)
	GetMatchesByPUUID(puuid string, numMatches int) (models.Matches, error)
	GetMatchByID(matchID string) (*match.Match, error)
}
