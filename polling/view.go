package polling

import (
	"github.com/bwmarrin/discordgo"
)

func view(poll Poll, session *discordgo.Session) {
	if poll.pollMessage != nil {
		session.ChannelMessageSendReply(poll.channel, "Here's the poll ^", poll.pollMessage.Reference())
	}
}
