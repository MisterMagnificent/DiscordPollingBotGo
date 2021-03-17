package polling

import (
	"fmt"
	"strings"

	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

var pollByChannel map[string]Poll = make(map[string]Poll)

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

	goBot.UpdateListeningStatus(config.BotPrefix + " for any commands")

	setup(&pollByChannel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func Cleanup() {
	fmt.Println("Cleanup called!")
	shutdown(goBot, pollByChannel)
	fmt.Println("Bot about to turn off!")
	goBot.Close()
	fmt.Println("Bot is turning off!")
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	var messCont = strings.ToLower(message.Content)
	if strings.HasPrefix(messCont, config.BotPrefix) {
		if message.Author.ID == BotID {
			return
		}

		if strings.HasPrefix(messCont, config.BotPrefix+"start") {

			pollByChannel[message.ChannelID] = start(session, message)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"add:") {

			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[message.ChannelID]
			addOption(&(temp), session, message)
			pollByChannel[message.ChannelID] = temp

		} else if messCont == config.BotPrefix+"bad" {

			bad(session, message)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"getrequests") {

			getFeatureList(session, message)

		} else if messCont == config.BotPrefix+"girl" {

			girl(session, message)

		} else if messCont == config.BotPrefix+"help" || messCont == config.BotPrefix+"elp" {

			help(session, message)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"remove:") {

			//Doesn't let you pass an address for some god forsaken reason, so temp variable workaround
			temp := pollByChannel[message.ChannelID]
			removeOption(&(temp), session, message)
			pollByChannel[message.ChannelID] = temp

		} else if messCont == config.BotPrefix+"repin" {

			pin(pollByChannel[message.ChannelID], session)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"request:") {

			addFeature(session, message)

		} else if strings.HasPrefix(messCont, config.BotPrefix+"reset") {

			var poll = pollByChannel[message.ChannelID]
			var newPoll = poll
			var split = strings.Split(messCont, " ")
			if poll.entries == nil || (len(split) > 1 && split[1] == "all") {
				newPoll = reset(poll, session, message)
			} else if len(split) == 1 {
				newPoll = resetCarryOver(poll, session, message, "")
			} else {
				newPoll = resetCarryOver(poll, session, message, split[1])
			}
			pollByChannel[message.ChannelID] = newPoll

		} else if strings.HasPrefix(messCont, config.BotPrefix+"result") {

			var poll = pollByChannel[message.ChannelID]
			var res = getResult(poll, session)
			_, _ = session.ChannelMessageSend(poll.channel, res)

		} else if messCont == config.BotPrefix+"runoff" {

			var poll = pollByChannel[message.ChannelID]
			runoff(poll, session, message)

		} else if messCont == config.BotPrefix+"runoffResult" {

			var poll = pollByChannel[message.ChannelID]
			var res = runoffRes(poll, session)
			_, _ = session.ChannelMessageSend(message.ChannelID, res)

		} else if messCont == config.BotPrefix+"shutdown" {
			if message.Author.ID == config.AdminID {
				session.ChannelMessageDelete(message.ChannelID, message.ID)
				forceShutdown(session, pollByChannel)
			}

		} else if messCont == config.BotPrefix+"view" {

			view(pollByChannel[message.ChannelID], session)

		} else {
			return
		}

		session.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}
