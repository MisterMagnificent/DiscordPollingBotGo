package polling

import ()

type Event struct {
	Description string
	NumPlayers  string
	Length      string
	Players     []string
	Times       map[string][]string
}

func NewEvent() Event {
	event := Event{Times: map[string][]string{}}
	return event
}
