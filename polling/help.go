package polling

import (
	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
)

func help(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Version 0.5.0 \n\nGithub: https://github.com/MisterMagnificent/DiscordPollingBotGo/tree/main/polling\n\nAll commands start with: '"+config.BotPrefix+"' \n'start: XXX' command to start a poll with XXX question, just start will add a default question, \n'add: XXX' command to add an option called XXX (Multiple separated by :, i.e. XXX:YYY:ZZZ), \n'result' to calculate winner, \n'runoff' to start a runoff poll between all tied winners.  'runoffResult' for that result.  Next reset must pass in winner to remove from carry over.\n'view' to get a link to the poll message\n'remove:' to remove an option from the poll\n'reset' to reset the poll and start a new one with all entries (that have more than "+string(config.MinCarryOver)+" votes) but the winner from the last one.   If you want a full reset, 'reset all'")
}
