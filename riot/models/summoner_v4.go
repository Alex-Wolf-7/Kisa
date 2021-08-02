package models

type SummonerV4Response struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int64  `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}
