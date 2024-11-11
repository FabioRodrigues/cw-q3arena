package main

import (
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
	loggerService := logger.NewLogger()
	loggerService.Info("Starting the app")
	sorterService := sorter2.NewSortService()
	parserService := parser.New()

	killSubscriber := subscribers.NewKillSubscriber()
	rankingSubscriber := subscribers.NewRankingSubscriber(sorterService)
	deathCauseSubscriber := subscribers.NewDeathCauseSubscriber()

	gameProcessor := gameprocessor.NewGameProcessor(loggerService, parserService, killSubscriber, rankingSubscriber, deathCauseSubscriber)

	gameLoader := loader.NewLoaderService(ioadapter.NewIOAdapter(), gameProcessor)
	result, err := gameLoader.Load("seed/seed.txt")
	if err != nil {
		loggerService.Error(err)
	}
	loggerService.Info("output:")
	fmt.Println(result)
}
