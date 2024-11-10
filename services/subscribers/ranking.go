package subscribers

import (
	"cw-q3arena/constants"
	"cw-q3arena/entities"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"encoding/json"
	"errors"
	"sync"
)

type RankingDatabaseModel struct {
	Ranking []reportmodels.RankingReport
	Kills   map[int]reportmodels.RankingReport
}

type RankingSubscriber struct {
	mux    *sync.Mutex
	db     map[string]RankingDatabaseModel
	sorter services.Sorter
}

func NewRankingSubscriber(sorter services.Sorter) *RankingSubscriber {
	return &RankingSubscriber{
		mux:    &sync.Mutex{},
		db:     map[string]RankingDatabaseModel{},
		sorter: sorter,
	}
}

func (s *RankingSubscriber) Receive(gameId string, data any) {
	killEvent := data.(entities.Kill)
	s.mux.Lock()
	defer s.mux.Unlock()
	db, found := s.db[gameId]
	if !found {
		db = RankingDatabaseModel{
			Ranking: []reportmodels.RankingReport{},
			Kills:   map[int]reportmodels.RankingReport{},
		}
	}

	victim, found := db.Kills[killEvent.VictimId]
	if !found {
		victim = reportmodels.RankingReport{
			PlayerId:   killEvent.VictimId,
			PlayerName: killEvent.VictimName,
			Kills:      0,
		}
		db.Kills[killEvent.VictimId] = victim
	}

	killer, found := db.Kills[killEvent.KillerId]
	if !found && killEvent.KillerName != constants.WorldUsername {
		killer = reportmodels.RankingReport{
			PlayerId:   killEvent.KillerId,
			PlayerName: killEvent.KillerName,
			Kills:      0,
		}
		db.Kills[killEvent.KillerId] = killer
	}

	// In case of world user, we just decrease kills from the victim
	if killEvent.KillerName == constants.WorldUsername {
		victim.Kills -= 1
		db.Kills[killEvent.VictimId] = victim
	} else {
		killer.Kills += 1
		db.Kills[killEvent.KillerId] = killer
	}

	items := []reportmodels.RankingReport{}

	for _, item := range db.Kills {
		items = append(items, item)
	}

	db.Ranking = s.sorter.SortRankings(items)
	s.db[gameId] = db

}

func (s *RankingSubscriber) GetReport(gameId string) ([]reportmodels.RankingReport, error) {
	db, found := s.db[gameId]
	if !found {
		return nil, errors.New("game not found")
	}

	ranking := []reportmodels.RankingReport{}
	for _, item := range db.Ranking {
		if item.Kills < 0 {
			item.Kills = 0
		}
		ranking = append(ranking, item)
	}

	return ranking, nil
}

func (s *RankingSubscriber) GetSerializedReport(gameId string) (string, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	report, err := s.GetReport(gameId)
	if err != nil {
		return "", err
	}

	result, err := json.Marshal(report)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
