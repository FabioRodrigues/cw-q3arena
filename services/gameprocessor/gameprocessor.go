package gameprocessor

import (
	"context"
	"cw-q3arena/events"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"sync"
)

type GameProcessor struct {
	loggerService        services.Logger
	parser               services.Parser
	killSubscriber       services.Subscriber
	rankingSubscriber    services.Subscriber
	deathCauseSubscriber services.Subscriber
	subscribers          map[events.EventType][]services.Subscriber
}

func NewGameProcessor(
	loggerService services.Logger,
	parser services.Parser,
	killSubscriber services.Subscriber,
	rankingSubscriber services.Subscriber,
	deathCauseSubscriber services.Subscriber) services.GameProcessor {
	return &GameProcessor{
		loggerService: loggerService,
		subscribers: map[events.EventType][]services.Subscriber{
			events.EventKill: {killSubscriber, rankingSubscriber, deathCauseSubscriber},
		},
		killSubscriber:       killSubscriber,
		rankingSubscriber:    rankingSubscriber,
		deathCauseSubscriber: deathCauseSubscriber,
		parser:               parser,
	}
}

// ProcessGame receives a full game to process
func (p GameProcessor) ProcessGame(ctx context.Context, gameId string, game []string) reportmodels.ProcessorReport {

	var wg sync.WaitGroup
	lineChan := make(chan string, 3)

	// We parallelize the event triggering by 3 works (no strong reason for 3 workers, just a number that came to my mind)
	for i := 0; i < 3; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case line, ok := <-lineChan:
					if !ok {
						return
					}

					event, data, err := p.parser.Parse(line)
					if err != nil {
						p.loggerService.Error(err)
						wg.Done() // Ensure wg.Done() is called even on error
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
			}
		}()
	}

	for _, line := range game {
		select {
		case <-ctx.Done():
			break
		default:
			wg.Add(1)
			lineChan <- line
		}
	}

	wg.Wait()
	close(lineChan)

	killReport, err := p.killSubscriber.GetData(gameId)
	if err != nil {
		p.loggerService.Info("no kill reports found for game ", gameId)

	}

	rankingReport, err := p.rankingSubscriber.GetData(gameId)
	if err != nil {
		p.loggerService.Info("no ranking reports found for game ", gameId)
	}

	deathCauseReport, err := p.deathCauseSubscriber.GetData(gameId)
	if err != nil {
		p.loggerService.Info("no death cause reports found for game ", gameId)
	}

	return reportmodels.ProcessorReport{
		Game:             gameId,
		KillReport:       killReport,
		RankinReport:     rankingReport,
		DeathCauseReport: deathCauseReport,
	}
}
