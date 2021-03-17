package polling

import (
	"github.com/bwmarrin/discordgo"
)

func view(poll Poll, session *discordgo.Session) {
	if poll.PollMessage != nil {
		session.ChannelMessageSendReply(poll.Channel, "Here's the poll ^", poll.PollMessage.Reference())
	}
}
