package services

import (
	"context"
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
	ProcessGame(ctx context.Context, gameId string, game []string) reportmodels.ProcessorReport
}

type LoaderService interface {
	Load(ctx context.Context, path string) (string, error)
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Subscriber interface {
	Receive(gameId string, data any)
	GetData(gameId string) (map[string]any, error)
}
