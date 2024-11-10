package events

// EventType represents different types of game events in Quake 3 Arena.
type EventType int

const (
	EventClientConnect EventType = iota
	EventClientUserinfoChanged
	EventClientDisconnect
	EventKill
	EventItem
	EventInitGame
	EventShutdownGame
	EventUnknown
)
