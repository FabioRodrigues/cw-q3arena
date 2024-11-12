package subscribers

import (
	"cw-q3arena/entities"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"errors"
	"sync"
)

type KillSubscriber struct {
	mux *sync.Mutex
	db  map[string]reportmodels.KillReport
}

func NewKillSubscriber() services.Subscriber {
	return &KillSubscriber{
		mux: &sync.Mutex{},
		db:  map[string]reportmodels.KillReport{},
	}
}

func (s *KillSubscriber) Receive(gameId string, data any) {
	killData := data.(entities.Kill)
	s.mux.Lock()
	defer s.mux.Unlock()
	report, found := s.db[gameId]

	if !found {
		report = reportmodels.KillReport{}
	}

	report.AddPlayers(map[int]string{
		killData.KillerId: killData.KillerName,
		killData.VictimId: killData.VictimName},
	)

	report.AddKill(killData.KillerId, killData.KillerName, killData.VictimId)

	s.db[gameId] = report
}

func (s *KillSubscriber) getReport(gameId string) (reportmodels.KillReport, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	report, found := s.db[gameId]
	if !found {
		return reportmodels.KillReport{}, errors.New("report not found")
	}

	report.Kills = report.GetKills()

	return report, nil

}

func (s *KillSubscriber) GetData(gameId string) (map[string]any, error) {
	report, err := s.getReport(gameId)
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		gameId: map[string]any{
			"kills":       report.Kills,
			"total_kills": report.TotalKills,
			"players":     report.Players,
		},
	}

	return data, nil
}
