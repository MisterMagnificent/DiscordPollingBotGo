package polling

import (
	"github.com/bwmarrin/discordgo"
)

func pin(poll Poll, session *discordgo.Session) {
	session.ChannelMessagePin(poll.channel, poll.pollMessage.ID)
}

func unpin(poll Poll, session *discordgo.Session) {
	session.ChannelMessageUnpin(poll.channel, poll.pollMessage.ID)
}
