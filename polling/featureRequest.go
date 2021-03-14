package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

var featureRequestList []string = []string{}

func addFeature(session *discordgo.Session, message *discordgo.MessageCreate) {
	var split = strings.Split(message.Content, ": ")

	if len(split) > 1 {
		featureRequestList = append(featureRequestList, split[1])
		_, _ = session.ChannelMessageSend(message.ChannelID, "Feature request noted, thank you")
	}
}
