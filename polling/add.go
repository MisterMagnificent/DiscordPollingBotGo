package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func addOption(poll *Poll, session *discordgo.Session, channelID string, content string) {
	index := strings.IndexByte(content, ' ')
	chars := []rune(content)
	option := string(chars[index+1:])

	if option != "" {
		var split = strings.Split(option, ";")

		if len(split) > 0 {
			for index := 0; index < len(split); index++ {
				addHelper(poll, session, channelID, strings.TrimSpace(split[index]))
			}
		}
	}
}

func addHelper(poll *Poll, session *discordgo.Session, channelID string, element string) {
	var emote string = (*poll).Emotes[poll.LastLetter] //Pull from dictionary
	(*poll).LastLetter++

	(*poll).Entries[emote] = element
	(*poll).EntriesReverse[element] = emote
	(*poll).PollMessage.Content = (*poll).PollMessage.Content + "\n" + element + ": " + emote + "\n"

	_, _ = session.ChannelMessageEdit(channelID, (*poll).PollMessage.ID, (*poll).PollMessage.Content)
	go session.MessageReactionAdd(channelID, (*poll).PollMessage.ID, emote)
}
