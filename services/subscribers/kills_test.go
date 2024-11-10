package subscribers

import (
	"cw-q3arena/entities"
	"cw-q3arena/reportmodels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKillsReport(t *testing.T) {
	t.Run("should add initial kill data", func(t *testing.T) {
		subscriber := NewKillSubscriber()

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:22",
			KillerName: "Isgalamido",
			VictimName: "Dono da bola",
			KillerId:   1,
			VictimId:   2,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetReport("game_1")

		assert.NoError(t, err)
		assert.Equal(t, 1, report.TotalKills)
		assert.Equal(t, 2, len(report.Players))
		assert.Equal(t, map[string]int{"Isgalamido": 1}, report.Kills)

	})

	t.Run("should accumulate game kill events", func(t *testing.T) {
		subscriber := NewKillSubscriber()

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
			Timestamp:  "2:25",
			KillerName: "Fabio",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetReport("game_1")

		assert.NoError(t, err)
		assert.Equal(t, 2, report.TotalKills)
		assert.Equal(t, 3, len(report.Players))
		assert.Equal(t, 1, report.Kills["Isgalamido"])
		assert.Equal(t, 1, report.Kills["Fabio"])

	})

	t.Run("should not add world as a player", func(t *testing.T) {
		subscriber := NewKillSubscriber()

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
			Timestamp:  "2:25",
			KillerName: "<world>",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetReport("game_1")

		assert.NoError(t, err)
		assert.Equal(t, 2, report.TotalKills)
		assert.Equal(t, 2, len(report.Players))
		// 0 because World has killed him. So -1 kill
		assert.Equal(t, map[string]int{"Isgalamido": 0}, report.Kills)

	})

	t.Run("should not mix game events", func(t *testing.T) {
		subscriber := NewKillSubscriber()

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:22",
			KillerName: "Isgalamido",
			VictimName: "Dono da bola",
			KillerId:   1,
			VictimId:   2,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		subscriber.Receive("game_2", entities.Kill{
			Timestamp:  "2:25",
			KillerName: "Fabio",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetReport("game_1")

		assert.NoError(t, err)
		assert.Equal(t, 1, report.TotalKills)
		assert.Equal(t, 2, len(report.Players))
		assert.Equal(t, map[string]int{"Isgalamido": 1}, report.Kills)

	})

	t.Run("should not return negative kills if world kills more than the player kills", func(t *testing.T) {
		subscriber := NewKillSubscriber()

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
			Timestamp:  "2:25",
			KillerName: "<world>",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})
		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:26",
			KillerName: "<world>",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetReport("game_1")

		assert.NoError(t, err)
		assert.Equal(t, 3, report.TotalKills)
		assert.Equal(t, 2, len(report.Players))
		assert.Equal(t, map[string]int{"Isgalamido": 0}, report.Kills)

	})

	t.Run("should return no data when no events received", func(t *testing.T) {
		subscriber := NewKillSubscriber()

		report, err := subscriber.GetReport("game_1")

		assert.Equal(t, reportmodels.KillReport{}, report)
		assert.Error(t, err)
		assert.Equal(t, "report not found", err.Error())

	})
}
