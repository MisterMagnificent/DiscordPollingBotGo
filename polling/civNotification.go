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

	var tag = ""
	switch name {
	case "Mr Magnificent":
		tag = "<@110933964654407680>"
	case "HaydenKale":
		tag = "<@427025678005960711>"
	case "KungFoody":
		tag = "<@604849954724511744>"
	case "Straker":
		tag = "<@289996211086163968>"
	case "ajar1189":
		tag = "<@179400614948634624>"
	case "Link":
		tag = "<@108799298849751040>"
	case "hcycyota":
		tag = "<@399914099011616768>"
	default:
		tag = name
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
