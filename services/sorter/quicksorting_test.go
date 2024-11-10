package sorter

import (
	"cw-q3arena/reportmodels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("should sort", func(t *testing.T) {
		rankings := []reportmodels.RankingReport{
			{"e", 5, 5},
			{"b", 2, 2},
			{"d", 4, 4},
			{"a", 1, 1},
			{"c", 3, 3},
		}

		sorted := NewSortService().SortRankings(rankings)

		assert.Equal(t, []reportmodels.RankingReport{
			{"e", 5, 5},
			{"d", 4, 4},
			{"c", 3, 3},
			{"b", 2, 2},
			{"a", 1, 1},
		}, sorted)
	})
}
