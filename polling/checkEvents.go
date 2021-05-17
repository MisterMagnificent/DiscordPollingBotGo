package polling

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strconv"
	"strings"
)

func checkEventsAfterReact(eventMan *EventManager, added *discordgo.MessageReaction, session *discordgo.Session) {
	userID := added.UserID
	chann, _ := session.UserChannelCreate(userID)
	if chann != nil {
		_, _ = session.ChannelMessageSend(chann.ID, "When are you available for "+(*eventMan).Events[added.Emoji.Name].Description+"?\nAnswer in one of the following formats: ('"+added.Emoji.Name+" XX/XX XX:XX Xh' OR '"+added.Emoji.Name+" ___day XX:XX Xh')")

		channelSet := channelsByUser[userID]
		if len(channelSet) == 0 {
			channelSet = make(map[string]struct{})
		}
		if _, ok := channelSet[eventMan.Channel]; !ok {
			channelSet[eventMan.Channel] = struct{}{}
			channelsByUser[userID] = channelSet
		}
	}
}

func respondedWithSchedule(eventMan *EventManager, author string, content string, channelID string, session *discordgo.Session) {
	inputStyle1Reg := regexp.MustCompile("^.*+ \\d\\d/\\d\\d \\d\\d:\\d\\d .*?h$")
	inputStyle2Reg := regexp.MustCompile("^.*+ .*+ \\d\\d:\\d\\d .*?h$")
	if inputStyle1Reg.MatchString(content) || inputStyle2Reg.MatchString(content) {

		emoteIndex := strings.Index(content, " ")
		emote := content[0:emoteIndex]

		//check if author voted for that emote
		var users, _ = session.MessageReactions((*eventMan).Channel, (*eventMan).EventMessage.ID, emote, 100, "", "")
		for _, user := range users {
			if user.ID == author {
				event, check := (*eventMan).Events[emote]
				if check {
					//Make sure time given is more than required time for event
					length, _ := strconv.ParseFloat(event.Length, 32)
					elements := strings.Split(content, " ")
					givenLen := elements[len(elements)-1]
					givenLenClean := givenLen[0 : len(givenLen)-2]
					givenLenNum, err := strconv.ParseFloat(givenLenClean, 32)

					if err == nil && givenLenNum >= length {
						event.Players = append(event.Players, author)
						event.Times[author] = append(event.Times[author], content[emoteIndex+1:])
						(*eventMan).Events[emote] = event
						_, _ = session.ChannelMessageSend(channelID, "Time received, feel free to keep adding times")
					} else {
						_, _ = session.ChannelMessageSend(channelID, "Length available needs to be longer than or equal to "+event.Length+" hours for this event")
					}
				}

				return
			}
		}
	} else {
		_, _ = session.ChannelMessageSend(channelID, "Input did not match valid formats")
	}
}
