package polling

import ()

type Event struct {
	Description string
	NumPlayers  string
	Length      string
}

func NewEvent() Event {
	event := Event{}
	return event
}
