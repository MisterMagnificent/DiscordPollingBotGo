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

func suslax(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://cdn.betterttv.net/emote/60d1f3a98ed8b373e4217ccd/3x.gif")
}

func susgreed(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://cdn.discordapp.com/attachments/110933673733271552/943425937196871690/susgreed.gif")
}

func jason(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://64.media.tumblr.com/1c1038af9688ef46ede96759bd661524/tumblr_ojvv31CcsE1rgpgw0o4_540.gif")
}

func elp(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "https://media1.tenor.com/images/347047c2bf923c8e0861bb76e8f2644b/tenor.gif")
}

func nft(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "Do not right click copy plz https://www.cheatsheet.com/wp-content/uploads/2020/09/DArcy-Carden-2.jpg")
}
