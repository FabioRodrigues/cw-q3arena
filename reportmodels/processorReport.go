package reportmodels

type ProcessorReport struct {
	Game             string
	KillReport       map[string]any
	RankinReport     map[string]any
	DeathCauseReport map[string]any
}
