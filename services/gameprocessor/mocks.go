package gameprocessor

import (
	"context"
	"cw-q3arena/reportmodels"
)

type Mock struct {
	ProcessGameFn func(ctx context.Context, gameId string, game []string) reportmodels.ProcessorReport
}

func (m Mock) ProcessGame(ctx context.Context, gameId string, game []string) reportmodels.ProcessorReport {
	return m.ProcessGameFn(ctx, gameId, game)
}
