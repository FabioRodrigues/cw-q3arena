package gameprocessor

import (
	"cw-q3arena/events"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"cw-q3arena/services/subscribers"
	"sync"
)

type GameProcessor struct {
	loggerService     services.Logger
	parser            services.Parser
	killSubscriber    subscribers.Subscriber
	rankingSubscriber subscribers.Subscriber
	subscribers       map[events.EventType][]subscribers.Subscriber
}

func NewGameProcessor(
	loggerService services.Logger,
	parser services.Parser,
	killSubscriber subscribers.Subscriber,
	rankingSubscriber subscribers.Subscriber) *GameProcessor {
	return &GameProcessor{
		loggerService: loggerService,
		subscribers: map[events.EventType][]subscribers.Subscriber{
			events.EventKill: {killSubscriber, rankingSubscriber},
		},
		killSubscriber:    killSubscriber,
		rankingSubscriber: rankingSubscriber,
		parser:            parser,
	}
}

func (p GameProcessor) ProcessGame(gameId string, game []string) reportmodels.ProcessorReport {

	var wg sync.WaitGroup
	lineChan := make(chan string, 3)

	for i := 0; i < 3; i++ {
		go func() {
			for line := range lineChan {
				event, data, err := p.parser.Parse(line)
				if err != nil {
					p.loggerService.Error(err)
					continue
				}

				// Process the parsed event
				subs, found := p.subscribers[event]
				if found {
					for _, subscriber := range subs {
						subscriber.Receive(gameId, data)
					}
				}
				wg.Done()
			}
		}()
	}

	for _, line := range game {
		wg.Add(1)
		lineChan <- line
	}

	wg.Wait()
	close(lineChan)

	killReport, err := p.killSubscriber.GetData(gameId)
	if err != nil {
		p.loggerService.Info("no kill reports found for game", gameId)

	}

	rankingReport, err := p.rankingSubscriber.GetData(gameId)
	if err != nil {
		p.loggerService.Info("no ranking reports found for game", gameId)
	}

	return reportmodels.ProcessorReport{
		Game:         gameId,
		KillReport:   killReport,
		RankinReport: rankingReport,
	}
}
