package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func start(session *discordgo.Session, message *discordgo.MessageCreate) Poll {
	var channel = message.ChannelID

	var poll Poll = New()
	poll.channel = channel

	var questionMessage = "Which one"
	var command = strings.Split(message.Content, ":")
	if len(command) > 1 {
		questionMessage = command[1]
	}
	poll.question = questionMessage

	poll.pollMessage, _ = session.ChannelMessageSend(channel, poll.question+":")

	//pin
	pin(poll, session, message)

	return poll
}
