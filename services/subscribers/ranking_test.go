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

		report, err := subscriber.GetReport("game_1")
		assert.NoError(t, err)
		assert.Equal(t, []reportmodels.RankingReport{
			{"Isgalamido", 1, 2},
			{"Dono da bola", 2, 1},
			{"Fabio", 3, 0},
		}, report)
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

		report, err := subscriber.GetReport("game_1")
		assert.NoError(t, err)
		assert.Len(t, report, 3)
		assert.Equal(t, reportmodels.RankingReport{"Isgalamido", 1, 2}, report[0])
		assert.Equal(t, 0, report[1].Kills)
		assert.Equal(t, 0, report[2].Kills)
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

		_, err := subscriber.GetReport("game_2")
		assert.Error(t, err)
		assert.Equal(t, "game not found", err.Error())
	})
}
