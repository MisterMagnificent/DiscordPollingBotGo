package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// %schedule Description of the event you're going for 2-4[# of ppl] #hoursLength[optional]
func scheduleEvent(session *discordgo.Session, channelID string, options string) {
	eventMan := eventManagerByChannel[channelID]
	if eventMan.LastEmote == 0 {
		eventManagerByChannel[channelID] = NewEventManager()
	}

	lastSpaceIndex := strings.LastIndex(options, " ")
	lastString := options[lastSpaceIndex+1:]
	restOfOptions := options[0:lastSpaceIndex]
	description := ""
	numPeople := ""
	hoursLength := ""

	if strings.Contains(lastString, "-") {
		numPeople = lastString
		description = restOfOptions
	} else {
		hoursLength = lastString
		numPeopleIndex := strings.LastIndex(restOfOptions, " ")
		numPeople = restOfOptions[numPeopleIndex+1:]
		description = restOfOptions[0:numPeopleIndex]

		if !strings.Contains(numPeople, "-") {
			return
		}
	}

	eve := NewEvent()
	eve.Description = description
	eve.NumPlayers = numPeople
	eve.Length = hoursLength

	var emote string = eventMan.Emotes[eventMan.LastEmote] //Pull from dictionary
	eventMan.LastEmote++

	eventMan.Events[emote] = eve
	eventMan.EventsReverse[eve] = emote
}
