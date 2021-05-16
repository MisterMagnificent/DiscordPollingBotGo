package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func checkEventsAfterReact(eventMan *EventManager, added *discordgo.MessageReaction, session *discordgo.Session) {
	userID := added.UserID
	chann, _ := session.UserChannelCreate(userID)
	if chann != nil {
		_, _ = session.ChannelMessageSend(chann.ID, "When are you available for "+(*eventMan).Events[added.Emoji.Name].Description+"?\nAnswer in one of the following formats: ('"+added.Emoji.Name+" XX/XX XX:XX Xh' OR '"+added.Emoji.Name+" XXXday XX:XX Xh')")

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
	emoteIndex := strings.Index(content, " ")
	emote := content[0:emoteIndex]

	//check if author voted for that emote
	var users, _ = session.MessageReactions((*eventMan).Channel, (*eventMan).EventMessage.ID, emote, 100, "", "")
	for _, user := range users {
		if user.ID == author {
			event, check := (*eventMan).Events[emote]
			if check {
				event.Players = append(event.Players, author)
				event.Times[author] = append(event.Times[author], content[emoteIndex+1:])
				(*eventMan).Events[emote] = event
				_, _ = session.ChannelMessageSend(channelID, "Time received, feel free to keep adding times")
			}

			return
		}
	}
}
