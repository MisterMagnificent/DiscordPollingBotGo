package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func addOption(poll *Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	var split = strings.Split(message.Content, ":")

	if len(split) > 1 {
		for index := 1; index < len(split); index++ {
			addHelper(poll, session, message, strings.TrimSpace(split[index]))
		}
	}
}

func addHelper(poll *Poll, session *discordgo.Session, message *discordgo.MessageCreate, element string) {
	var emote string = (*poll).emotes[poll.lastLetter] //Pull from dictionary
	(*poll).lastLetter++

	(*poll).entries[emote] = element
	(*poll).entriesReverse[element] = emote
	(*poll).pollMessage.Content = (*poll).pollMessage.Content + "\n" + element + ": " + emote + "\n"

	_, _ = session.ChannelMessageEdit(message.ChannelID, (*poll).pollMessage.ID, (*poll).pollMessage.Content)
	go session.MessageReactionAdd(message.ChannelID, (*poll).pollMessage.ID, emote)
}
