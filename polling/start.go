package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func start(session *discordgo.Session, message *discordgo.MessageCreate) Poll {
	var channel = message.ChannelID

	var poll Poll = New()
	poll.Channel = channel

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
