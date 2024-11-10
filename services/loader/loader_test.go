package loader

import (
	"cw-q3arena/infra/ioadapter"
	"cw-q3arena/services/gameprocessor"
	"cw-q3arena/services/parser"
	sorter2 "cw-q3arena/services/sorter"
	"cw-q3arena/services/subscribers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoader(t *testing.T) {
	t.Run("should load file", func(t *testing.T) {
		sorter := sorter2.NewSortService()
		parser := parser.New()

		killSubscriber := subscribers.NewKillSubscriber()
		rankingSubscriber := subscribers.NewRankingSubscriber(sorter)

		gameProcessor := gameprocessor.NewGameProcessor(parser, killSubscriber, rankingSubscriber)

		svc := NewLoaderService(ioadapter.NewIOAdapter(), gameProcessor)
		result, err := svc.Load("../../seed/seed.txt")
		assert.NoError(t, err)
		fmt.Println(result)
	})
}
