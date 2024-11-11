package sorter

import (
	"cw-q3arena/reportmodels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("should sort", func(t *testing.T) {
		rankings := []reportmodels.RankingReport{
			{PlayerName: "e", PlayerId: 5, Kills: 5},
			{PlayerName: "b", PlayerId: 2, Kills: 2},
			{PlayerName: "d", PlayerId: 4, Kills: 4},
			{PlayerName: "a", PlayerId: 1, Kills: 1},
			{PlayerName: "c", PlayerId: 3, Kills: 3},
		}

		sorted := NewSortService().SortRankings(rankings)

		assert.Equal(t, []reportmodels.RankingReport{
			{PlayerName: "e", PlayerId: 5, Kills: 5},
			{PlayerName: "d", PlayerId: 4, Kills: 4},
			{PlayerName: "c", PlayerId: 3, Kills: 3},
			{PlayerName: "b", PlayerId: 2, Kills: 2},
			{PlayerName: "a", PlayerId: 1, Kills: 1},
		}, sorted)
	})
}
