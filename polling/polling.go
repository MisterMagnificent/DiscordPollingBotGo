package polling

import (
	"fmt"
	"strings"

	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
	"github.com/prprprus/scheduler"
)

var BotID string
var goBot **discordgo.Session
var ourScheduler **scheduler.Scheduler

var pollByChannel map[string]Poll = make(map[string]Poll)
var eventManagerByChannel map[string]EventManager = make(map[string]EventManager)

var channelsByUser map[string]map[string]struct{} = make(map[string]map[string]struct{})

func Start() {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goBot = &bot

	tempSched, err := scheduler.NewScheduler(1000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ourScheduler = &tempSched

	u, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	bot.AddHandler(messageHandler)
	bot.AddHandler(updateHandler)
	bot.AddHandler(reactHandler)

	err = bot.Open()

	bot.UpdateListeningStatus(config.BotPrefix + " for any commands")

	setup(&pollByChannel, &eventManagerByChannel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func Cleanup() {
	fmt.Println("Cleanup called!")
	shutdown(*goBot, pollByChannel)
}

func reactHandler(session *discordgo.Session, message *discordgo.MessageReactionAdd) {
	if message.MessageReaction != nil && message.MessageReaction.ChannelID != "" {
		if message.MessageReaction.UserID == BotID {
			return
		}

		temp := eventManagerByChannel[message.MessageReaction.ChannelID]
		//checkEventsAfterReact(&(temp), message.MessageReaction, session)
		eventManagerByChannel[message.MessageReaction.ChannelID] = temp
	} else {
		fmt.Println("Issues with this React: %s", message.MessageReaction)
	}
}

func updateHandler(session *discordgo.Session, message *discordgo.MessageUpdate) {
	if message.Message != nil && message.Message.Author != nil {
		parseCommand(session, message.Message.ID, message.Message.Content, message.Message.ChannelID, message.Message.Author)
	} else {
		fmt.Println("Issues with this message: %s", message.Message)
	}
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	parseCommand(session, message.ID, message.Content, message.ChannelID, message.Author)
}

func parseCommand(session *discordgo.Session, id string, content string, channelID string, author *discordgo.User) {
	var authorID = author.ID
	if authorID == BotID {
		return
	}

	var messCont = strings.ToLower(content)
	if strings.HasPrefix(messCont, config.BotPrefix) {

		// Get the command (pre first space), then grab everything after as the "options" as they could be a single option with spacing, that parsing is left to the command itself
		index := strings.IndexByte(messCont, ' ')
		lowerChars := []rune(messCont)
		command := ""
		options := ""

		if index != -1 {
			command = string(lowerChars[1:index])

			//Options is the only thing that needs to save capitalization
			chars := []rune(content)
			options = string(chars[index+1:])
		} else {
			command = string(lowerChars[1:])
		}

		switch command {
		case "start":
			if _, ok := pollByChannel[channelID]; ok {
				_, _ = session.ChannelMessageSend(channelID, "A poll already exists for this channel.  If you want to force a new one, use '***start!***'")
			} else {
				pollByChannel[channelID] = start(session, channelID, options, pollByChannel)
			}
		case "start!":
			pollByChannel[channelID] = start(session, channelID, options, pollByChannel)
		case "add":
			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[channelID]
			addOption(&(temp), session, channelID, options, author)
			pollByChannel[channelID] = temp
		case "bad":
			bad(session, channelID)
		case "derek":
			derek(session, channelID)
		case "elp":
			elp(session, channelID)
		case "getrequests":
			getFeatureList(session, channelID)
		case "girl":
			girl(session, channelID)
		case "help":
			help(session, channelID, options)
		case "jason":
			jason(session, channelID)
		case "remove":
			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[channelID]
			removeOption(&(temp), session, options)
			pollByChannel[channelID] = temp
		case "repin":
			pin(pollByChannel[channelID], session)
		case "request":
			addFeature(session, channelID, options)
		case "reset":
			var poll = pollByChannel[channelID]
			var newPoll = poll
			if poll.Entries == nil || options == "all" {
				newPoll = reset(pollByChannel, session, channelID)
			} else {
				newPoll = resetCarryOver(poll, session, options)
			}
			pollByChannel[channelID] = newPoll
		case "result":
			var poll = pollByChannel[channelID]
			var res = getResult(poll, session)
			_, _ = session.ChannelMessageSend(poll.Channel, res)
		case "runoff":
			temp := pollByChannel[channelID]
			runoff(&(temp), session)
			pollByChannel[channelID] = temp
		case "runoffresult":
			var poll = pollByChannel[channelID]
			var res = runoffRes(poll, session)
			_, _ = session.ChannelMessageSend(channelID, res)
		case "schedule":
			eventManager := eventManagerByChannel[channelID]
			if eventManager.LastEmote == 0 {
				eventManager = NewEventManager()
			}
			scheduleEvent(&(eventManager), session, channelID, options)
			eventManagerByChannel[channelID] = eventManager
		case "schedulemessage":
			scheduleMessage(session, channelID, options)
		case "shutdown":
			if authorID == config.AdminID {
				forceShutdown(session, pollByChannel)
			}
		case "view":
			view(pollByChannel[channelID], session)
		default:
			return
		}

		session.ChannelMessageDelete(channelID, id)
	} else {
		for chann, _ := range channelsByUser[authorID] {
			eventManager := eventManagerByChannel[chann]
			if eventManager.LastEmote != 0 {
				respondedWithSchedule(&(eventManager), authorID, content, channelID, session)
				eventManagerByChannel[chann] = eventManager
			}
		}
	}
}
