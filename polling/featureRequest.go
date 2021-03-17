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
		_, _ = session.ChannelMessageSend(message.ChannelID, "Feature request of '"+split[1]+"' noted, thank you")
	}
}

func getFeatureList(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Features currently requested: \n"+strings.Join(featureRequestList, "\n"))
}
