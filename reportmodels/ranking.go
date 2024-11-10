package reportmodels

type RankingReport struct {
	PlayerName string `json:"player_name"`
	PlayerId   int    `json:"player_id"`
	Kills      int    `json:"kills"`
}
