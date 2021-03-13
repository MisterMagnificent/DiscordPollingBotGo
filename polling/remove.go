package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func removeOption(poll *Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	var split = strings.Split(message.Content, ": ")

	if len(split) > 1 {
		removeOptionHelper(poll, session, split[1], true)
	}
}
func removeOptionHelper(poll *Poll, session *discordgo.Session, element string, removeText bool) {
	var tempEmote = (*poll).entriesReverse[element]
	delete(poll.entriesReverse, element)
	delete((*poll).entries, tempEmote)
	(*poll).entries[tempEmote] = tempEmote

	//Emotes forever grows here, but too much of a pain in the ass for fixing the edge case where removes happen infinitely on a forever resetting poll
	(*poll).emotes[len((*poll).emotes)] = tempEmote

	if removeText {
		(*poll).pollMessage.Content = strings.ReplaceAll((*poll).pollMessage.Content, "\n"+element+": "+tempEmote+"\n", "")

		_, _ = session.ChannelMessageEdit(poll.channel, (*poll).pollMessage.ID, (*poll).pollMessage.Content)
		go session.MessageReactionsRemoveEmoji(poll.channel, (*poll).pollMessage.ID, tempEmote)
	}
}
