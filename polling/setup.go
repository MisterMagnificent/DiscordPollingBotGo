package polling

import (
	"encoding/json"
	"io/ioutil"
)

func setup(pollByChannel *map[string]Poll, eventByChannel *map[string]EventManager) {
	//Load from file here
	var file, _ = ioutil.ReadFile("polls.json")
	_ = json.Unmarshal(file, pollByChannel)
}
