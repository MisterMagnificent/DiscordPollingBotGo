package polling

import (
	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type EventManager struct {
	Channel      string
	EventMessage *discordgo.Message
	Emotes       map[int]string
	Events       map[string]Event
	LastEmote    int
}

func NewEventManager() EventManager {
	var emotes map[int]string = map[int]string{}
	var emoteList = strings.Split(config.Emotes, ",")
	for index, element := range emoteList {
		emotes[index] = element
	}

	eventMan := EventManager{LastEmote: 0, Emotes: emotes, Events: map[string]Event{}}
	return eventMan
}
