package polling

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

// %schedule Description of the event you're going for 2-4[# of ppl] #hoursLength[optional]
func scheduleEvent(eventMan *EventManager, session *discordgo.Session, channelID string, options string) {
	inputStyleReg := regexp.MustCompile("^.+ \\d+?-\\d+? \\d*\\.*\\d+?$")
	if inputStyleReg.MatchString(options) {
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

		var emote string = (*eventMan).Emotes[(*eventMan).LastEmote] //Pull from dictionary
		(*eventMan).LastEmote++

		((*eventMan).Events)[emote] = eve
		(*eventMan).Channel = channelID

		eveMess := emote + ": " + eve.Description + " - " + eve.NumPlayers + " Players - for " + eve.Length + " hour(s)"

		if (*eventMan).EventMessage == nil {
			(*eventMan).EventMessage, _ = session.ChannelMessageSend(channelID, "Events\nVote for which you'd like to join:\n"+eveMess)
		} else {
			(*eventMan).EventMessage.Content = (*eventMan).EventMessage.Content + "\n\n" + eveMess
			_, _ = session.ChannelMessageEdit(channelID, (*eventMan).EventMessage.ID, (*eventMan).EventMessage.Content)
		}
		go session.MessageReactionAdd(channelID, (*eventMan).EventMessage.ID, emote)
	} else {
		session.ChannelMessageSend(channelID, "The schedule message does not match format of '___ X-X X'")
	}
}
