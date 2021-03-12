package polling

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotID string
var goBot *discordgo.Session
var pollMessageByChannel map[string]*discordgo.Message = make(map[string]*discordgo.Message)
var runoffMessageByChannel map[string]*discordgo.Message = make(map[string]*discordgo.Message)
var lastLetterByChannel map[string]int = make(map[string]int)
var emotes map[int]string = map[int]string{
	0:  "🇦",
	1:  "🇧",
	2:  "🇨",
	3:  "🇩",
	4:  "🇪",
	5:  "🇫",
	6:  "🇬",
	7:  "🇭",
	8:  "🇮",
	9:  "🇯",
	10: "🇰",
	11: "🇱",
	12: "🇲",
	13: "🇳",
	14: "🇴",
	15: "🇵",
	16: "🇶",
	17: "🇷",
	18: "🇸",
	19: "🇹",
	20: "🇺",
	21: "🇻",
	22: "🇼",
	23: "🇽",
	24: "🇾",
	25: "🇿",
}
var entriesByChannel map[string]map[string]string = make(map[string]map[string]string)
var entriesReverseByChannel map[string]map[string]string = make(map[string]map[string]string)
var entries map[string]string = map[string]string{
	"🇦": "🇦",
	"🇧": "🇧",
	"🇨": "🇨",
	"🇩": "🇩",
	"🇪": "🇪",
	"🇫": "🇫",
	"🇬": "🇬",
	"🇭": "🇭",
	"🇮": "🇮",
	"🇯": "🇯",
	"🇰": "🇰",
	"🇱": "🇱",
	"🇲": "🇲",
	"🇳": "🇳",
	"🇴": "🇴",
	"🇵": "🇵",
	"🇶": "🇶",
	"🇷": "🇷",
	"🇸": "🇸",
	"🇹": "🇹",
	"🇺": "🇺",
	"🇻": "🇻",
	"🇼": "🇼",
	"🇽": "🇽",
	"🇾": "🇾",
	"🇿": "🇿",
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"add:") {
			//Add option to poll
			addOption(s, m)
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"result") {
			var res = getResult(s, m)
			_, _ = s.ChannelMessageSend(m.ChannelID, res)
			//unpin
			s.ChannelMessageUnpin(m.ChannelID, pollMessageByChannel[m.ChannelID].Content)
		}

		if m.Content == config.BotPrefix+"repin" {
			//Pin again
			s.ChannelMessagePin(m.ChannelID, pollMessageByChannel[m.ChannelID].Content)
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"reset") {
			var split = strings.Split(m.Content, " ")
			if entriesByChannel[m.ChannelID] == nil {
				reset(s, m)
			} else if len(split) == 1 || split[1] == "" {
				resetCarryOver(s, m, "")
			} else if len(split) > 1 && split[1] == "all" {
				reset(s, m)
			} else {
				resetCarryOver(s, m, split[1])
			}
		}

		if m.Content == config.BotPrefix+"runoff" {
			runoff(s, m)
		}

		if m.Content == config.BotPrefix+"runoffResult" {
			var res = runoffRes(s, m)
			_, _ = s.ChannelMessageSend(m.ChannelID, res)
		}

		if m.Content == config.BotPrefix+"help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "All commands start with: '"+config.BotPrefix+"' \n'add: XXX' command to add an option called XXX, \n'result' to calculate winner, \n'runoff' to start a runoff poll between all tied winners.  'runoffResult' for that result.  Next reset must pass in winner to remove from carry over.\n'reset' to reset the poll and start a new one with all but the winner from the last one.   If you want a full reset, 'reset all'")
		}
	}
}

func addOption(s *discordgo.Session, message *discordgo.MessageCreate) {
	var emote string = emotes[lastLetterByChannel[message.ChannelID]] //Pull from dictionary
	var split = strings.Split(message.Content, ": ")
	fmt.Println(entriesByChannel)
	entriesByChannel[message.ChannelID][emote] = split[1]
	entriesReverseByChannel[message.ChannelID][split[1]] = emote
	pollMessageByChannel[message.ChannelID].Content = pollMessageByChannel[message.ChannelID].Content + "\n" + split[1] + ": " + emote + "\n"
	_, _ = s.ChannelMessageEdit(message.ChannelID, pollMessageByChannel[message.ChannelID].ID, pollMessageByChannel[message.ChannelID].Content)
	go s.MessageReactionAdd(message.ChannelID, pollMessageByChannel[message.ChannelID].ID, emote)
	lastLetterByChannel[message.ChannelID]++
}

func getResult(s *discordgo.Session, message *discordgo.MessageCreate) string {
	return getResultHelper(s, message, pollMessageByChannel[message.ChannelID])
}

func getResultHelper(s *discordgo.Session, message *discordgo.MessageCreate, reactionMessage *discordgo.Message) string {
	var biggest int = 0
	var emote string
	for key, _ := range entriesByChannel[message.ChannelID] {
		var users, _ = s.MessageReactions(message.ChannelID, reactionMessage.ID, key, 100, "", "")
		var size = len(users)
		if size >= biggest {
			if size > biggest {
				biggest = size
				emote = key
			} else {
				emote = emote + "," + key
			}
		}
	}

	var split = strings.Split(emote, ",")
	var entry = entriesByChannel[message.ChannelID][split[0]]
	var result = entry + " (" + entriesReverseByChannel[message.ChannelID][entry] + ")"
	for i := 1; i < len(split); i++ {
		entry = entriesByChannel[message.ChannelID][split[i]]
		result = result + ", " + entry + " (" + entriesReverseByChannel[message.ChannelID][entry] + ")"
	}
	return result
}

func copyMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}

//remove the winner
func resetCarryOver(s *discordgo.Session, message *discordgo.MessageCreate, winner string) {
	var won string = winner
	if winner == "" {
		//find winner
		var res string = getResult(s, message)
		var splitRes []string = strings.Split(res, " (")
		won = splitRes[0]
	}
	lastLetterByChannel[message.ChannelID]--
	var emoteKey string = entriesReverseByChannel[message.ChannelID][won]
	var lastLetter string = emotes[lastLetterByChannel[message.ChannelID]]
	var newWord string = entriesByChannel[message.ChannelID][lastLetter]
	entriesByChannel[message.ChannelID][emoteKey] = newWord
	entriesByChannel[message.ChannelID][lastLetter] = lastLetter
	entriesReverseByChannel[message.ChannelID][newWord] = emoteKey
	delete(entriesReverseByChannel[message.ChannelID], won)

	var newMessage string = "Poll reset.  New poll with carryover has begun:"
	pollMessageByChannel[message.ChannelID], _ = s.ChannelMessageSend(message.ChannelID, newMessage)
	for key, val := range entriesByChannel[message.ChannelID] {
		if entries[val] == "" {
			pollMessageByChannel[message.ChannelID].Content = pollMessageByChannel[message.ChannelID].Content + "\n" + val + ": " + key + "\n"
			_, _ = s.ChannelMessageEdit(message.ChannelID, pollMessageByChannel[message.ChannelID].ID, pollMessageByChannel[message.ChannelID].Content)
			go s.MessageReactionAdd(message.ChannelID, pollMessageByChannel[message.ChannelID].ID, key)
		}
	}
	//Pin
	s.ChannelMessagePin(message.ChannelID, pollMessageByChannel[message.ChannelID].Content)
}

func reset(s *discordgo.Session, message *discordgo.MessageCreate) {
	lastLetterByChannel[message.ChannelID] = 0
	entriesByChannel[message.ChannelID] = copyMap(entries)
	entriesReverseByChannel[message.ChannelID] = copyMap(entries)
	runoffMessageByChannel[message.ChannelID] = nil

	pollMessageByChannel[message.ChannelID], _ = s.ChannelMessageSend(message.ChannelID, "Poll reset.  New poll has begun:")
	//Pin
	s.ChannelMessagePin(message.ChannelID, pollMessageByChannel[message.ChannelID].Content)
}

func runoff(s *discordgo.Session, message *discordgo.MessageCreate) {
	var split []string = strings.Split(getResult(s, message), ", ")
	var messageOutput string = "Runoff poll:\n"
	var emotes []string
	for _, key := range split {
		var splitEmote []string = strings.Split(key, "(")
		var splitEmoteFin = strings.Split(splitEmote[1], ")")
		emotes = append(emotes, splitEmoteFin[0])

		messageOutput += key + "\n"
	}
	//for each result, add one
	runoffMessageByChannel[message.ChannelID], _ = s.ChannelMessageSend(message.ChannelID, messageOutput)

	for _, emote := range emotes {
		go s.MessageReactionAdd(message.ChannelID, runoffMessageByChannel[message.ChannelID].ID, emote)
	}
}

func runoffRes(s *discordgo.Session, message *discordgo.MessageCreate) string {
	if runoffMessageByChannel[message.ChannelID] != nil {
		return getResultHelper(s, message, runoffMessageByChannel[message.ChannelID])
	}
	return ""
}
