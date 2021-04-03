package polling

import (
	"github.com/bwmarrin/discordgo"
)

func view(poll Poll, session *discordgo.Session) {
	if poll.PollMessage != nil {
		if val, ok := poll.LastMessage["view"]; ok {
			session.ChannelMessageDelete(poll.Channel, val.id)
		}
		poll.LastMessage["view"], _ = session.ChannelMessageSendReply(poll.Channel, "Here's the poll ^", poll.PollMessage.Reference())
	}
}
