package subscribers

import (
	"cw-q3arena/entities"
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

		report, err := subscriber.GetData("game_1")

		assert.NoError(t, err)
		totalKills, kills, players := getReportData(report, "game_1")

		assert.Equal(t, 1, totalKills)
		assert.Equal(t, 2, len(players))
		assert.Equal(t, map[string]int{"Isgalamido": 1}, kills)

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

		report, err := subscriber.GetData("game_1")
		totalKills, kills, players := getReportData(report, "game_1")

		assert.NoError(t, err)
		assert.Equal(t, 2, totalKills)
		assert.Equal(t, 3, len(players))
		assert.Equal(t, 1, kills["Isgalamido"])
		assert.Equal(t, 1, kills["Fabio"])

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

		report, err := subscriber.GetData("game_1")
		totalKills, kills, players := getReportData(report, "game_1")

		assert.NoError(t, err)
		assert.Equal(t, 2, totalKills)
		assert.Equal(t, 2, len(players))
		// 0 because World has killed him. So -1 kill
		assert.Equal(t, map[string]int{"Isgalamido": 0}, kills)

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

		report, err := subscriber.GetData("game_1")
		totalKills, kills, players := getReportData(report, "game_1")

		assert.NoError(t, err)
		assert.Equal(t, 1, totalKills)
		assert.Equal(t, 2, len(players))
		assert.Equal(t, map[string]int{"Isgalamido": 1}, kills)

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

		report, err := subscriber.GetData("game_1")
		totalKills, kills, players := getReportData(report, "game_1")

		assert.NoError(t, err)
		assert.Equal(t, 3, totalKills)
		assert.Equal(t, 2, len(players))
		assert.Equal(t, map[string]int{"Isgalamido": 0}, kills)

	})

	t.Run("should return no data when no events received", func(t *testing.T) {
		subscriber := NewKillSubscriber()

		report, err := subscriber.GetData("game_1")

		assert.Len(t, report, 0)
		assert.Error(t, err)
		assert.Equal(t, "report not found", err.Error())

	})
}

func getReportData(report map[string]any, gameId string) (int, map[string]int, []string) {
	result := report[gameId].(map[string]interface{})
	totalKills := result["total_kills"].(int)
	kills := result["kills"].(map[string]int)
	players := result["players"].([]string)

	return totalKills, kills, players
}
