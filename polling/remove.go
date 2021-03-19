package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func removeOption(poll *Poll, session *discordgo.Session, content string) {
	index := strings.IndexByte(content, ' ')
	chars := []rune(content)
	option := string(chars[index+1:])

	if option != "" {
		removeOptionHelper(poll, session, option, true)
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
