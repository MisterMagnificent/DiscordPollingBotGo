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

func elp(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://media1.tenor.com/images/347047c2bf923c8e0861bb76e8f2644b/tenor.gif")
}
