package polling

import (
	"../config"
	"github.com/bwmarrin/discordgo"
)

func help(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Version 0.3.0 \n\nGithub: https://github.com/MisterMagnificent/DiscordPollingBotGo/tree/main/polling\nAll commands start with: '"+config.BotPrefix+"' \n'start: XXX' command to start a poll with XXX question, just start will add a default question, \n'add: XXX' command to add an option called XXX, \n'result' to calculate winner, \n'runoff' to start a runoff poll between all tied winners.  'runoffResult' for that result.  Next reset must pass in winner to remove from carry over.\n'reset' to reset the poll and start a new one with all but the winner from the last one.   If you want a full reset, 'reset all'")
}
