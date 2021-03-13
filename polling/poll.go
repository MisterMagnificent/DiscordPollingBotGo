package polling

import (
	"../config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Poll struct {
	channel        string
	question       string
	pollMessage    *discordgo.Message
	runoffMessage  *discordgo.Message
	lastLetter     int
	emotes         map[int]string
	entries        map[string]string
	entriesReverse map[string]string
	isMovie        bool
}

func New() Poll {
	var emotes map[int]string = map[int]string{}
	var entries map[string]string = map[string]string{}
	var emoteList = strings.Split(config.Emotes, ",")
	for index, element := range emoteList {
		emotes[index] = element
		entries[element] = element
	}

	poll := Poll{lastLetter: 0, emotes: emotes, entries: entries, entriesReverse: map[string]string{}}
	//create emotes
	return poll
}

func copyMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
func copyIntMap(m map[int]string) map[int]string {
	cp := make(map[int]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
