package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

var featureRequestList []string = []string{}

func addFeature(session *discordgo.Session, channelID string, option string) {
	if option != "" {
		featureRequestList = append(featureRequestList, option)
		_, _ = session.ChannelMessageSend(channelID, "Feature request of '"+option+"' noted, thank you")
	}
}

func getFeatureList(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "Features currently requested: \n"+strings.Join(featureRequestList, "\n"))
}
