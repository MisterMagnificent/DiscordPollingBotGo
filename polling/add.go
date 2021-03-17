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
	var emote string = (*poll).Emotes[poll.LastLetter] //Pull from dictionary
	(*poll).LastLetter++

	(*poll).Entries[emote] = element
	(*poll).EntriesReverse[element] = emote
	(*poll).PollMessage.Content = (*poll).PollMessage.Content + "\n" + element + ": " + emote + "\n"

	_, _ = session.ChannelMessageEdit(message.ChannelID, (*poll).PollMessage.ID, (*poll).PollMessage.Content)
	go session.MessageReactionAdd(message.ChannelID, (*poll).PollMessage.ID, emote)
}
