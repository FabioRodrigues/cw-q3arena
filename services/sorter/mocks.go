package sorter

import "cw-q3arena/reportmodels"

type Mock struct {
	SortRankingsFn func(rankings []reportmodels.RankingReport) []reportmodels.RankingReport
}

func (m Mock) SortRankings(rankings []reportmodels.RankingReport) []reportmodels.RankingReport {
	if m.SortRankingsFn != nil {
		return m.SortRankingsFn(rankings)
	}
	return []reportmodels.RankingReport{}
}
