package polling

import (
	"github.com/bwmarrin/discordgo"
)

func bad(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "What up forknuts?")
}

func girl(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "I'm not a girl")
}

func derek(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://media.giphy.com/media/xUOxf02RS2TWDj0xtC/giphy.gif")
}
