package polling

import (
	"github.com/bwmarrin/discordgo"
)

func bad(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "What up forknuts?")
}
