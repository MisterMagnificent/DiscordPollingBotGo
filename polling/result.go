package polling

import (
	"github.com/bwmarrin/discordgo"
)

func getResult(poll Poll, session *discordgo.Session) string {
	return getResultHelper(poll, session)
}

func getResultHelper(poll Poll, session *discordgo.Session) string {
	var biggest int = 0
	var emote []string
	for key, _ := range poll.entries {
		var users, _ = session.MessageReactions(poll.channel, poll.pollMessage.ID, key, 100, "", "")
		var size = len(users)
		if size >= biggest {
			if size > biggest {
				biggest = size
				emote = []string{key}
			} else {
				emote = append(emote, key)
			}
		}
	}

	result = appendNamesToEmotes(poll, emote)
	return result
}

func appendNamesToEmotes(poll Poll, emoteList []string) string {
	var entry = poll.entries[emoteList[0]]
	var result = entry + " (" + poll.entriesReverse[entry] + ")"
	for i := 1; i < len(emoteList); i++ {
		entry = poll.entries[emoteList[i]]
		result = result + ", " + entry + " (" + poll.entriesReverse[entry] + ")"
	}
}
