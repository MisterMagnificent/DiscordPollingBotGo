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

	err = bot.Open()

	bot.UpdateListeningStatus(config.BotPrefix + " for any commands")

	setup(&pollByChannel)
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

func updateHandler(session *discordgo.Session, message *discordgo.MessageUpdate) {
	parseCommand(session, message.Message.ID, message.Message.Content, message.Message.ChannelID, message.Message.Author.ID)
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	parseCommand(session, message.ID, message.Content, message.ChannelID, message.Author.ID)
}

func parseCommand(session *discordgo.Session, id string, content string, channelID string, authorID string) {
	var messCont = strings.ToLower(content)
	if strings.HasPrefix(messCont, config.BotPrefix) {
		if authorID == BotID {
			return
		}

		if strings.HasPrefix(messCont, config.BotPrefix+"start") {

			pollByChannel[channelID] = start(session, channelID, content, pollByChannel)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"add ") {

			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[channelID]
			addOption(&(temp), session, channelID, content)
			pollByChannel[channelID] = temp

		} else if messCont == config.BotPrefix+"bad" {

			bad(session, channelID)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"getrequests") {

			getFeatureList(session, channelID)

		} else if messCont == config.BotPrefix+"girl" {

			girl(session, channelID)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"help") || strings.HasPrefix(messCont, config.BotPrefix+"elp") {

			help(session, channelID, content)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"remove ") {

			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[channelID]
			removeOption(&(temp), session, content)
			pollByChannel[channelID] = temp

		} else if messCont == config.BotPrefix+"repin" {

			pin(pollByChannel[channelID], session)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"request ") {

			addFeature(session, channelID, content)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"reset") {

			var poll = pollByChannel[channelID]
			var newPoll = poll
			var split = strings.Split(messCont, " ")
			if poll.Entries == nil || (len(split) > 1 && split[1] == "all") {
				newPoll = reset(pollByChannel, session, channelID, content)
			} else if len(split) == 1 {
				newPoll = resetCarryOver(poll, session, "")
			} else {
				newPoll = resetCarryOver(poll, session, split[1])
			}
			pollByChannel[channelID] = newPoll

		} else if strings.HasPrefix(messCont, config.BotPrefix+"result") {

			var poll = pollByChannel[channelID]
			var res = getResult(poll, session)
			_, _ = session.ChannelMessageSend(poll.Channel, res)

		} else if messCont == config.BotPrefix+"runoff" {

			var poll = pollByChannel[channelID]
			runoff(poll, session)

		} else if messCont == config.BotPrefix+"runoffResult" {

			var poll = pollByChannel[channelID]
			var res = runoffRes(poll, session)
			_, _ = session.ChannelMessageSend(channelID, res)

		} else if messCont == config.BotPrefix+"schedule" {
			schedule(session, channelID, content)
		} else if messCont == config.BotPrefix+"shutdown" {
			if authorID == config.AdminID {
				forceShutdown(session, pollByChannel)
			}

		} else if messCont == config.BotPrefix+"view" {

			view(pollByChannel[channelID], session)

		} else {
			return
		}

		session.ChannelMessageDelete(channelID, id)
	}
}
