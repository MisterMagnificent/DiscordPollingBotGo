package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
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

	var result = appendNamesToEmotes(poll, emote)
	return result
}

func appendNamesToEmotes(poll Poll, emoteList []string) string {
	var result = ""
	for i := 0; i < len(emoteList); i++ {
		var entry = poll.entries[emoteList[i]]
		var item = poll.entriesReverse[entry]
		if i > 0 {
			result = result + ", "
		}
		result = result + entry + " (" + item

		if poll.isMovie {
			result = result + " : https://www.justwatch.com/us/movie/" + strings.ReplaceAll(entry, " ", "-")
		}
		result = result + ")"
	}
	return result
}
