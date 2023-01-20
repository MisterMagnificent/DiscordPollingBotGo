package polling

import ()

func civNotification(content) {
	indexAt := strings.IndexByte(content, '@')
	indexHash := strings.IndexByte(content, '#')
	allChars := []rune(content)
	var name = string(allChars[indexAt+1 : indexHash-19])

	var restString = string(allChars[indexHash:])
	var firstSpace = strings.IndexByte(restString, ' ')
	var restAllChars = []rune(restString)
	var turn = string(restAllChars[1:firstSpace])
	var gameName = string(restAllChars[firstSpace+4 : len(restAllChars)-1])

	var tag = name
	value, exists := config.tags[name]
	if exists {
		tag = value
	}

	if tag != "" {
		allowArray := [](discordgo.AllowedMentionType){}
		allowArray = append(allowArray, discordgo.AllowedMentionTypeUsers)
		var messageSend = &discordgo.MessageSend{
			Content: "You're up, " + tag + " in game: " + gameName + ": turn " + turn,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: allowArray,
			},
		}
		_, _ = session.ChannelMessageSendComplex(channelID, messageSend)
		session.ChannelMessageDelete(channelID, id)
	}
}
