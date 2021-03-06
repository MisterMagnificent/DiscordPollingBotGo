package polling

import (
	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

//remove the winner
func resetCarryOver(poll Poll, session *discordgo.Session, winner string) Poll {
	unpin(poll, session)
	var won string = winner
	if winner == "" {
		//find winner
		var res string = getResult(poll, session)
		var splitRes []string = strings.Split(res, " (")
		won = splitRes[0]
	}

	removeOptionHelper(&(poll), session, won, false)

	var oldMessageId = poll.PollMessage.ID
	var newMessage string = "Poll reset.  New poll with carryover has begun:"
	poll.PollMessage, _ = session.ChannelMessageSend(poll.Channel, newMessage)

	for key, val := range poll.EntriesReverse {
		var users, _ = session.MessageReactions(poll.Channel, oldMessageId, val, 100, "", "")
		var size = len(users)

		if size >= config.MinCarryOver {
			poll.PollMessage.Content = poll.PollMessage.Content + "\n" + key + ": " + val + "\n"
			_, _ = session.ChannelMessageEdit(poll.Channel, poll.PollMessage.ID, poll.PollMessage.Content)
			go session.MessageReactionAdd(poll.Channel, poll.PollMessage.ID, val)
		}
	}
	//Pin
	pin(poll, session)
	return poll
}

func reset(pollByChannel map[string]Poll, session *discordgo.Session, channelID string) Poll {
	return start(session, channelID, pollByChannel[channelID].Question, pollByChannel)
}
