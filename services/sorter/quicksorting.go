package sorter

import (
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
)

func NewSortService() services.Sorter {
	return Sorter{}
}

type Sorter struct{}

// SortRankings is a very simple implementation of a quicksort for sorting the rankings within a game (desc order)
func (s Sorter) SortRankings(rankings []reportmodels.RankingReport) []reportmodels.RankingReport {
	if len(rankings) < 2 {
		return rankings
	}

	orderedRankings := []reportmodels.RankingReport{}

	picked := rankings[len(rankings)-1]
	left := []reportmodels.RankingReport{}
	right := []reportmodels.RankingReport{}

	for _, ranking := range rankings {
		if ranking.PlayerId == picked.PlayerId {
			continue
		}

		if ranking.Kills >= picked.Kills {
			left = append(left, ranking)
			continue
		}
		right = append(right, ranking)
	}

	orderedRankings = append(orderedRankings, s.SortRankings(left)...)
	orderedRankings = append(orderedRankings, picked)
	orderedRankings = append(orderedRankings, s.SortRankings(right)...)
	return orderedRankings
}
