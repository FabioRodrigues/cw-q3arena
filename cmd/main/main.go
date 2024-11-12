package main

import (
	"context"
	"cw-q3arena/infra/ioadapter"
	"cw-q3arena/services/gameprocessor"
	"cw-q3arena/services/loader"
	"cw-q3arena/services/logger"
	"cw-q3arena/services/parser"
	sorter2 "cw-q3arena/services/sorter"
	"cw-q3arena/services/subscribers"
	"fmt"
)

// Entrypoint of the application
func main() {
	// initializing logger
	loggerService := logger.NewLogger()
	loggerService.Info("Starting the app")

	// initializing services
	sorterService := sorter2.NewSortService()
	parserService := parser.New()

	killSubscriber := subscribers.NewKillSubscriber()
	rankingSubscriber := subscribers.NewRankingSubscriber(sorterService)
	deathCauseSubscriber := subscribers.NewDeathCauseSubscriber()

	ioAdapter := ioadapter.NewIOAdapter()

	gameProcessor := gameprocessor.NewGameProcessor(loggerService, parserService, killSubscriber, rankingSubscriber, deathCauseSubscriber)

	gameLoader := loader.NewLoaderService(ioAdapter, gameProcessor)
	result, err := gameLoader.Load(context.Background(), "seed/seed.txt")
	if err != nil {
		loggerService.Error(err)
	}
	loggerService.Info("output:")
	fmt.Println(result)
}
