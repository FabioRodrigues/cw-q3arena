package services

import (
	"cw-q3arena/events"
	"cw-q3arena/reportmodels"
)

type Parser interface {
	Parse(line string) (events.EventType, any, error)
}

type Sorter interface {
	SortRankings(rankings []reportmodels.RankingReport) []reportmodels.RankingReport
}

type GameProcessor interface {
	ProcessGame(gameId string, game []string) reportmodels.ProcessorReport
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}
