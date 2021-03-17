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
	var tempEmote = (*poll).EntriesReverse[element]
	delete((*poll).EntriesReverse, element)
	delete((*poll).Entries, tempEmote)
	(*poll).Entries[tempEmote] = tempEmote

	//Emotes forever grows here, but too much of a pain in the ass for fixing the edge case where removes happen infinitely on a forever resetting poll
	(*poll).Emotes[len((*poll).Emotes)] = tempEmote

	if removeText {
		(*poll).PollMessage.Content = strings.ReplaceAll((*poll).PollMessage.Content, "\n"+element+": "+tempEmote+"\n", "")

		_, _ = session.ChannelMessageEdit(poll.Channel, (*poll).PollMessage.ID, (*poll).PollMessage.Content)
		go session.MessageReactionsRemoveEmoji(poll.Channel, (*poll).PollMessage.ID, tempEmote)
	}
}
