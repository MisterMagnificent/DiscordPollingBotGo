package polling

import (
	"github.com/bwmarrin/discordgo"
)

func pin(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessagePin(message.ChannelID, poll.pollMessage.Content)
}

func unpin(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageUnpin(message.ChannelID, poll.pollMessage.Content)
}
