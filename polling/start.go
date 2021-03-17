package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func start(session *discordgo.Session, message *discordgo.MessageCreate, pollByChannel map[string]Poll) Poll {
	var channel = message.ChannelID

	var poll Poll = New()
	poll.Channel = channel

	if val, ok := pollByChannel[poll.Channel]; ok {
		unpin(val, session)
	}

	var questionMessage = "Which one"
	var command = strings.Split(message.Content, ":")
	if len(command) > 1 {
		questionMessage = command[1]
		if strings.Contains(questionMessage, "movie") || strings.Contains(questionMessage, "Movie") {
			poll.IsMovie = true
		}
	}
	poll.Question = questionMessage

	poll.PollMessage, _ = session.ChannelMessageSend(channel, poll.Question+":")

	//pin
	pin(poll, session)

	return poll
}
