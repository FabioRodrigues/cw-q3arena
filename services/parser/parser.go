package parser

import (
	"cw-q3arena/constants"
	"cw-q3arena/entities"
	"cw-q3arena/events"
	"cw-q3arena/services"
	"errors"
	"regexp"
	"strconv"
)

func New() services.Parser {
	return Parser{
		killRegex: regexp.MustCompile(constants.KillPattern),
	}
}

type Parser struct {
	killRegex *regexp.Regexp
}

func (p Parser) Parse(line string) (events.EventType, any, error) {
	if event, result, err := p.tryParseKill(line); err == nil {
		return event, result, err
	}

	return events.EventUnknown, nil, errors.New("unknown event")
}

func (p Parser) tryParseKill(line string) (events.EventType, any, error) {
	if p.killRegex.Match([]byte(line)) {
		matches := p.killRegex.FindStringSubmatch(line)
		if len(matches) != 8 {
			return events.EventKill, nil, errors.New("couldn't parse kill event")
		}
		killerId, err := strconv.Atoi(matches[2])
		if err != nil {
			killerId = 0
		}
		victimId, err := strconv.Atoi(matches[3])
		if err != nil {
			victimId = 0
		}
		methodId, err := strconv.Atoi(matches[4])
		if err != nil {
			methodId = 0
		}

		return events.EventKill, entities.Kill{
			Timestamp:  matches[1],
			KillerId:   killerId,
			VictimId:   victimId,
			MethodId:   methodId,
			KillerName: matches[5],
			VictimName: matches[6],
			MethodName: matches[7],
		}, nil
	}
	return events.EventUnknown, nil, errors.New("couldn't parse kill event")
}
