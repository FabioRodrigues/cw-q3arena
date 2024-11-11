package subscribers

import (
	"cw-q3arena/entities"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services/sorter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRankingSubscriber(t *testing.T) {
	t.Run("should get report ordered", func(t *testing.T) {
		subscriber := NewRankingSubscriber(sorter.NewSortService())
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:22",
			KillerName: "Isgalamido",
			VictimName: "Dono da bola",
			KillerId:   1,
			VictimId:   2,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:23",
			KillerName: "Isgalamido",
			VictimName: "Fabio",
			KillerId:   1,
			VictimId:   3,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:24",
			KillerName: "Dono da bola",
			VictimName: "Fabio",
			KillerId:   2,
			VictimId:   3,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetData("game_1")
		ranking := report["game_1"].([]reportmodels.RankingReport)

		assert.NoError(t, err)
		assert.Equal(t, []reportmodels.RankingReport{
			{PlayerName: "Isgalamido", PlayerId: 1, Kills: 2},
			{PlayerName: "Dono da bola", PlayerId: 2, Kills: 1},
			{PlayerName: "Fabio", PlayerId: 3},
		}, ranking)
	})

	t.Run("should get report ordered with world player", func(t *testing.T) {
		subscriber := NewRankingSubscriber(sorter.NewSortService())
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:22",
			KillerName: "Isgalamido",
			VictimName: "Dono da bola",
			KillerId:   1,
			VictimId:   2,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:23",
			KillerName: "Isgalamido",
			VictimName: "Fabio",
			KillerId:   1,
			VictimId:   3,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:24",
			KillerName: "Dono da bola",
			VictimName: "Fabio",
			KillerId:   2,
			VictimId:   3,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:25",
			KillerName: "<world>",
			VictimName: "Dono da bola",
			KillerId:   4,
			VictimId:   2,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetData("game_1")
		ranking := report["game_1"].([]reportmodels.RankingReport)

		assert.NoError(t, err)
		assert.Len(t, ranking, 3)
		assert.Equal(t, reportmodels.RankingReport{PlayerName: "Isgalamido", PlayerId: 1, Kills: 2}, ranking[0])
		assert.Equal(t, 0, ranking[1].Kills)
		assert.Equal(t, 0, ranking[2].Kills)
	})

	t.Run("should return error when no game found", func(t *testing.T) {
		subscriber := NewRankingSubscriber(sorter.NewSortService())
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:24",
			KillerName: "Dono da bola",
			VictimName: "Fabio",
			KillerId:   2,
			VictimId:   3,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		_, err := subscriber.GetData("game_2")
		assert.Error(t, err)
		assert.Equal(t, "game not found", err.Error())
	})
}
