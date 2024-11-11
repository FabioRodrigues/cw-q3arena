package subscribers

import (
	"cw-q3arena/entities"
	"errors"
	"sync"
)

type DeathCauseSubscriber struct {
	mux *sync.Mutex
	db  map[string]map[string]int
}

func NewDeathCauseSubscriber() Subscriber {
	return &DeathCauseSubscriber{
		mux: &sync.Mutex{},
		db:  map[string]map[string]int{},
	}
}

func (s *DeathCauseSubscriber) Receive(gameId string, data any) {
	killEvent := data.(entities.Kill)
	s.mux.Lock()
	defer s.mux.Unlock()
	db, found := s.db[gameId]
	if !found {
		db = map[string]int{}
	}

	causeCount, found := db[killEvent.MethodName]
	if !found {
		causeCount = 0
	}

	db[killEvent.MethodName] = causeCount + 1
	s.db[gameId] = db
}

func (s DeathCauseSubscriber) GetData(gameId string) (map[string]any, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	game, found := s.db[gameId]
	if !found {
		return nil, errors.New("game not found")
	}

	result := map[string]any{}
	for k, v := range game {
		result[k] = v
	}

	return result, nil
}
