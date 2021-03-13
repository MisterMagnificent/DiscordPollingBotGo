package polling

import (
	"../config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

//remove the winner
func resetCarryOver(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate, winner string) Poll {
	unpin(poll, session)
	var won string = winner
	if winner == "" {
		//find winner
		var res string = getResult(poll, session)
		var splitRes []string = strings.Split(res, " (")
		won = splitRes[0]
	}

	removeOptionHelper(&(poll), session, won, false)

	var oldMessageId = poll.pollMessage.ID
	var newMessage string = "Poll reset.  New poll with carryover has begun:"
	poll.pollMessage, _ = session.ChannelMessageSend(poll.channel, newMessage)

	for key, val := range poll.entries {
		var users, _ = session.MessageReactions(poll.channel, oldMessageId, val, 100, "", "")
		var size = len(users)

		if poll.entries[val] == "" && size > config.MinCarryOver {
			poll.pollMessage.Content = poll.pollMessage.Content + "\n" + val + ": " + key + "\n"
			_, _ = session.ChannelMessageEdit(poll.channel, poll.pollMessage.ID, poll.pollMessage.Content)
			go session.MessageReactionAdd(poll.channel, poll.pollMessage.ID, key)
		}
	}
	//Pin
	pin(poll, session)
	return poll
}

func reset(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate) Poll {
	unpin(poll, session)
	return start(session, message)
}
