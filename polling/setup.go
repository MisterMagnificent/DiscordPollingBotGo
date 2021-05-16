package polling

import (
	"encoding/json"
	"io/ioutil"
)

func setup(pollByChannel *map[string]Poll, eventByChannel *map[string]EventManager) {
	//Load from file here
	var pollFile, _ = ioutil.ReadFile("polls.json")
	_ = json.Unmarshal(pollFile, pollByChannel)

	var eventFile, _ = ioutil.ReadFile("events.json")
	_ = json.Unmarshal(eventFile, eventByChannel)
}
