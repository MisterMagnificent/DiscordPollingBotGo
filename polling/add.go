package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func addOption(poll *Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	var emote string = (*poll).emotes[poll.lastLetter] //Pull from dictionary
	(*poll).lastLetter++

	var split = strings.Split(message.Content, ": ")
	(*poll).entries[emote] = split[1]
	(*poll).entriesReverse[split[1]] = emote
	(*poll).pollMessage.Content = (*poll).pollMessage.Content + "\n" + split[1] + ": " + emote + "\n"

	_, _ = session.ChannelMessageEdit(message.ChannelID, (*poll).pollMessage.ID, (*poll).pollMessage.Content)
	go session.MessageReactionAdd(message.ChannelID, (*poll).pollMessage.ID, emote)
}
