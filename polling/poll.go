package polling

import (
	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Poll struct {
	Channel        string
	Question       string
	LastMessage    map[string]*discordgo.Message
	PollMessage    *discordgo.Message
	RunoffMessage  *discordgo.Message
	LastLetter     int
	Emotes         map[int]string
	Entries        map[string]string
	EntriesReverse map[string]string
	IsMovie        bool
}

func NewPoll() Poll {
	var emotes map[int]string = map[int]string{}
	var entries map[string]string = map[string]string{}
	var emoteList = strings.Split(config.Emotes, ",")
	for index, element := range emoteList {
		emotes[index] = element
		entries[element] = element
	}

	poll := Poll{LastLetter: 0, Emotes: emotes, Entries: entries, EntriesReverse: map[string]string{}, LastMessage: map[string]*discordgo.Message{}}
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
