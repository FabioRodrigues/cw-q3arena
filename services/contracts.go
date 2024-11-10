package services

import "cw-q3arena/events"

type Parser interface {
	Parse(line string) (events.EventType, any, error)
}
