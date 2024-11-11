package loader

import (
	"cw-q3arena/infra/ioadapter"
	"cw-q3arena/services/gameprocessor"
	"cw-q3arena/services/logger"
	"cw-q3arena/services/parser"
	sorter2 "cw-q3arena/services/sorter"
	"cw-q3arena/services/subscribers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoader(t *testing.T) {
	t.Run("e2e - should load file and get the report", func(t *testing.T) {
		// This is an end-to-end test that gets the
		sorterService := sorter2.NewSortService()
		parserService := parser.New()
		loggerService := logger.NewLogger()

		killSubscriber := subscribers.NewKillSubscriber()
		rankingSubscriber := subscribers.NewRankingSubscriber(sorterService)

		gameProcessor := gameprocessor.NewGameProcessor(loggerService, parserService, killSubscriber, rankingSubscriber)

		svc := NewLoaderService(ioadapter.NewIOAdapter(), gameProcessor)
		result, err := svc.Load("../../seed/seed_test.txt")
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}
