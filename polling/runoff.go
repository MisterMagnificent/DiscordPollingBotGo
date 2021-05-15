package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func runoff(poll *Poll, session *discordgo.Session) {
	var split []string = strings.Split(getResult(*poll, session), ", ")
	var messageOutput string = "Runoff poll:\n"
	var emotes []string
	for _, key := range split {
		var splitEmote []string = strings.Split(key, "(")
		var splitEmoteFin = strings.Split(splitEmote[1], ")")
		emotes = append(emotes, splitEmoteFin[0])

		messageOutput += key + "\n"
	}
	//for each result, add one
	(*poll).RunoffMessage, _ = session.ChannelMessageSend((*poll).Channel, messageOutput)

	for _, emote := range emotes {
		go session.MessageReactionAdd((*poll).Channel, (*poll).RunoffMessage.ID, emote)
	}
}

func runoffRes(poll Poll, session *discordgo.Session) string {
	if poll.RunoffMessage != nil {
		return getResultHelper(poll, poll.RunoffMessage, session)
	}
	return ""
}
