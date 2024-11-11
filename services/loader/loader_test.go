package loader

import (
	"bytes"
	"cw-q3arena/infra/ioadapter"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services/gameprocessor"
	"cw-q3arena/services/logger"
	"cw-q3arena/services/parser"
	sorter2 "cw-q3arena/services/sorter"
	"cw-q3arena/services/subscribers"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestLoader(t *testing.T) {
	lines := []string{
		" 12:24 Kill: 3 4 6: Isgalamido killed Zeh by MOD_ROCKET\n",
		" 12:25 Kill: 3 4 6: Isgalamido killed Zeh by MOD_ROCKET\n",
		" 12:26 Kill: 3 4 6: Isgalamido killed Zeh by MOD_ROCKET\n",
		" 12:13 InitGame: \n"}

	t.Run("e2e - should load file and get the report", func(t *testing.T) {
		// This is an end-to-end test that gets the file and process the report
		sorterService := sorter2.NewSortService()
		parserService := parser.New()
		loggerService := logger.NewLogger()

		killSubscriber := subscribers.NewKillSubscriber()
		rankingSubscriber := subscribers.NewRankingSubscriber(sorterService)
		deathCauseSubscriber := subscribers.NewDeathCauseSubscriber()

		gameProcessor := gameprocessor.NewGameProcessor(loggerService, parserService, killSubscriber, rankingSubscriber, deathCauseSubscriber)

		svc := NewLoaderService(ioadapter.NewIOAdapter(), gameProcessor)
		result, err := svc.Load("../../seed/seed_test.txt")
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("should return the report correctly", func(t *testing.T) {
		processGameCalled := false

		gameProcessor := gameprocessor.Mock{
			ProcessGameFn: func(gameId string, game []string) reportmodels.ProcessorReport {
				processGameCalled = true
				return reportmodels.ProcessorReport{
					Game: gameId,
					KillReport: map[string]any{
						gameId: map[string]any{
							"players": []string{"Isgalamido", "Zeh"},
						},
					},
				}
			},
		}

		index := 0
		mockIo := ioadapter.Mock{
			OpenFn: func(name string) (io.ReadCloser, error) {
				return &ioadapter.MockReadCloser{
					ReadFn: func(p []byte) (n int, err error) {
						if index > len(lines)-1 {
							return 0, io.EOF
						}
						reader := bytes.NewReader([]byte(lines[index]))
						index += 1
						return reader.Read(p)
					},
				}, nil
			},
		}
		svc := NewLoaderService(mockIo, gameProcessor)
		result, err := svc.Load("")
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.True(t, processGameCalled)
		var output []map[string]any
		err = json.Unmarshal([]byte(result), &output)
		assert.NoError(t, err)

		game := output[0]["game_1"].(map[string]interface{})
		killData := game["kill_data"].(map[string]interface{})
		players := killData["players"].([]interface{})
		assert.Len(t, players, 2)
	})
}
