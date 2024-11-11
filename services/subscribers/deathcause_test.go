package subscribers

import (
	"cw-q3arena/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeathCause(t *testing.T) {
	t.Run("should return report", func(t *testing.T) {
		subscriber := NewDeathCauseSubscriber()

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:25",
			KillerName: "Zeh",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:26",
			KillerName: "Zeh",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_ROCKET_SPLASH",
			MethodId:   10,
		})

		report, err := subscriber.GetData("game_1")
		assert.NoError(t, err)
		assert.Equal(t, 1, report["MOD_RAILGUN"])
		assert.Equal(t, 1, report["MOD_ROCKET_SPLASH"])
	})

	t.Run("should return error if no game found", func(t *testing.T) {
		subscriber := NewDeathCauseSubscriber()

		subscriber.Receive("game_1", entities.Kill{
			Timestamp:  "2:25",
			KillerName: "Zeh",
			VictimName: "Isgalamido",
			KillerId:   3,
			VictimId:   1,
			MethodName: "MOD_RAILGUN",
			MethodId:   10,
		})

		report, err := subscriber.GetData("game_2")
		assert.Error(t, err)
		assert.Nil(t, report)
		assert.ErrorContains(t, err, "game not found")
	})
}
