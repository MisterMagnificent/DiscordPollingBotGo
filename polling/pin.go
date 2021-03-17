package polling

import (
	"github.com/bwmarrin/discordgo"
)

func pin(poll Poll, session *discordgo.Session) {
	session.ChannelMessagePin(poll.Channel, poll.PollMessage.ID)
}

func unpin(poll Poll, session *discordgo.Session) {
	session.ChannelMessageUnpin(poll.Channel, poll.PollMessage.ID)
}
