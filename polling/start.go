package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func start(session *discordgo.Session, channel string, content string, pollByChannel map[string]Poll) Poll {
	var poll Poll = New()
	poll.Channel = channel

	if val, ok := pollByChannel[poll.Channel]; ok {
		unpin(val, session)
	}

	var questionMessage = "Which one"
	index := strings.IndexByte(content, ' ')
	chars := []rune(content)
	question := string(chars[index+1:])
	if question != "" {
		questionMessage = question
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
