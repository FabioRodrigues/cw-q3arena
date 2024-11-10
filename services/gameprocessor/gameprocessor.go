package gameprocessor

import (
	"cw-q3arena/events"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"cw-q3arena/services/subscribers"
	"fmt"
	"sync"
)

type GameProcessor struct {
	parser            services.Parser
	killSubscriber    subscribers.Subscriber
	rankingSubscriber subscribers.Subscriber
	subscribers       map[events.EventType][]subscribers.Subscriber
}

func NewGameProcessor(
	parser services.Parser,
	killSubscriber subscribers.Subscriber,
	rankingSubscriber subscribers.Subscriber) *GameProcessor {
	return &GameProcessor{
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
					fmt.Println("Error parsing line", line, err)
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

	killReport, err := p.killSubscriber.GetSerializedReport(gameId)
	if err != nil {
		fmt.Println("Error getting report", err)

	}

	rankingReport, err := p.rankingSubscriber.GetSerializedReport(gameId)
	if err != nil {
		fmt.Println("Error getting report", err)
	}

	return reportmodels.ProcessorReport{
		Game:         gameId,
		KillReport:   killReport,
		RankinReport: rankingReport,
	}
}
