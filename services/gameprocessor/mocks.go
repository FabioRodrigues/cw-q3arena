package gameprocessor

import "cw-q3arena/reportmodels"

type Mock struct {
	ProcessGameFn func(gameId string, game []string) reportmodels.ProcessorReport
}

func (m Mock) ProcessGame(gameId string, game []string) reportmodels.ProcessorReport {
	return m.ProcessGameFn(gameId, game)
}
