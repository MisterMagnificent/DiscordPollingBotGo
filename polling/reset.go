package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

//remove the winner
func resetCarryOver(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate, winner string) Poll {
	var won string = winner
	if winner == "" {
		//find winner
		var res string = getResult(poll, session)
		var splitRes []string = strings.Split(res, " (")
		won = splitRes[0]
	}

	poll.lastLetter--
	var lastLetterLookup = poll.lastLetter
	var emoteKey string = poll.entriesReverse[won]
	var lastLetter string = poll.emotes[lastLetterLookup]
	var newWord string = poll.entries[lastLetter]

	poll.entries[emoteKey] = newWord
	poll.entries[lastLetter] = lastLetter
	poll.entriesReverse[newWord] = emoteKey
	delete(poll.entriesReverse, won)

	var newMessage string = "Poll reset.  New poll with carryover has begun:"
	poll.pollMessage, _ = session.ChannelMessageSend(poll.channel, newMessage)
	for key, val := range poll.entries {
		if poll.entries[val] == "" {
			poll.pollMessage.Content = poll.pollMessage.Content + "\n" + val + ": " + key + "\n"
			_, _ = session.ChannelMessageEdit(poll.channel, poll.pollMessage.ID, poll.pollMessage.Content)
			go session.MessageReactionAdd(poll.channel, poll.pollMessage.ID, key)
		}
	}
	//Pin
	pin(poll, session, message)
	return poll
}

func reset(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate) Poll {
	poll.lastLetter = 0
	poll.entries = copyMap(entries)
	poll.entriesReverse = copyMap(entries)
	poll.runoffMessage = nil

	poll.pollMessage, _ = session.ChannelMessageSend(poll.channel, "Poll reset.  New poll has begun:")

	pin(poll, session, message)
	return poll
}
