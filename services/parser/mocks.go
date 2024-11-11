package parser

import "cw-q3arena/events"

type Mock struct {
	ParseFn func(line string) (events.EventType, any, error)
}

func (m Mock) Parse(line string) (events.EventType, any, error) {
	if m.ParseFn != nil {
		return m.ParseFn(line)
	}
	return events.EventUnknown, nil, nil
}
