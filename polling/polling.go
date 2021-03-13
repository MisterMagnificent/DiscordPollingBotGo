package polling

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotID string
var goBot *discordgo.Session

var pollByChannel map[string]Poll = make(map[string]Poll)

var pollMessageByChannel map[string]*discordgo.Message = make(map[string]*discordgo.Message)
var runoffMessageByChannel map[string]*discordgo.Message = make(map[string]*discordgo.Message)
var lastLetterByChannel map[string]int = make(map[string]int)
var entriesByChannel map[string]map[string]string = make(map[string]map[string]string)
var entriesReverseByChannel map[string]map[string]string = make(map[string]map[string]string)

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

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	if strings.HasPrefix(message.Content, config.BotPrefix) {
		if message.Author.ID == BotID {
			return
		}

		if strings.HasPrefix(message.Content, config.BotPrefix+"start") {
			pollByChannel[message.ChannelID] = start(session, message)
		}

		if strings.HasPrefix(message.Content, config.BotPrefix+"add:") {
			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[message.ChannelID]
			addOption(&(temp), session, message)
			pollByChannel[message.ChannelID] = temp
		}

		if message.Content == config.BotPrefix+"repin" {
			pin(pollByChannel[message.ChannelID], session, message)
		}

		if message.Content == config.BotPrefix+"help" {
			help(session, message)
		}

		if strings.HasPrefix(message.Content, config.BotPrefix+"result") {
			var poll = pollByChannel[message.ChannelID]
			var res = getResult(poll, session)
			_, _ = session.ChannelMessageSend(poll.channel, res)
			//unpin
			unpin(poll, session, message)
		}

		//--------------------------Migrated-------------------------------------

		if strings.HasPrefix(message.Content, config.BotPrefix+"reset") {
			var poll = pollByChannel[message.ChannelID]
			var split = strings.Split(message.Content, " ")
			if entriesByChannel[message.ChannelID] == nil {
				reset(session, message)
			} else if len(split) == 1 || split[1] == "" {
				resetCarryOver(poll, session, message, "")
			} else if len(split) > 1 && split[1] == "all" {
				reset(session, message)
			} else {
				resetCarryOver(poll, session, message, split[1])
			}
		}

		if message.Content == config.BotPrefix+"runoff" {
			var poll = pollByChannel[message.ChannelID]
			runoff(poll, session, message)
		}

		if message.Content == config.BotPrefix+"runoffResult" {
			var poll = pollByChannel[message.ChannelID]
			var res = runoffRes(poll, session, message)
			_, _ = session.ChannelMessageSend(message.ChannelID, res)
		}
	}
}

//remove the winner
func resetCarryOver(poll Poll, s *discordgo.Session, message *discordgo.MessageCreate, winner string) {
	var won string = winner
	if winner == "" {
		//find winner
		var res string = getResult(poll, s)
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

func runoff(poll Poll, s *discordgo.Session, message *discordgo.MessageCreate) {
	var split []string = strings.Split(getResult(poll, s), ", ")
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

func runoffRes(poll Poll, s *discordgo.Session, message *discordgo.MessageCreate) string {
	if runoffMessageByChannel[message.ChannelID] != nil {
		return getResultHelper(poll, s)
	}
	return ""
}
