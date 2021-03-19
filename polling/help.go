package polling

import (
	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func help(session *discordgo.Session, channelID string, content string) {
	index := strings.IndexByte(content, ' ')

	if index == -1 {
		carry := strconv.Itoa(config.MinCarryOver)
		_, _ = session.ChannelMessageSend(channelID, "Version 1.0.0 Beta \n\nHi, here to help with whatever you need! I can make one live poll per channel at a time.\nCOMMANDS:\n\nAll commands start with: '"+config.BotPrefix+"' \n'***start XXX***' command to start a poll with XXX question, just start will add a default question, \n'***add XXX***' command to add an option called XXX (Multiple separated by ;, i.e. XXX;YYY;ZZZ), \n'***result***' to calculate winner, \n'***view***' to get a link to the poll message\n'***reset***' to reset the poll and start a new one with all entries (that have more than "+carry+" votes) but the winner from the last one.   If you want a full reset, '***reset all***'\n\n '***help advanced***' for more commands! https://64.media.tumblr.com/daf862b49b82e49a47354b14c5143363/tumblr_oefwev7pLA1vvi3bvo8_250.gif")
	} else {
		chars := []rune(content)
		option := string(chars[index+1:])
		if option == "advanced" {
			helpAdvanced(session, channelID)
		}
	}

}

func helpAdvanced(session *discordgo.Session, channelID string) {
	_, _ = session.ChannelMessageSend(channelID, "Version 1.0.0 Beta \n\n'***schedule***' to schedule a 'view' call automatically, format dayOfWeek:Hour:Minute.  Day Of Week starting with Sunday as 0 up to Saturday as 6.  Hour is UTC timezone\n'***runoff***' to start a runoff poll between all tied winners.  '***runoffResult***' for that result.  Next reset must pass in winner to remove from carry over.\n'***remove***' to remove an option from the poll\n\nGithub: https://github.com/MisterMagnificent/DiscordPollingBotGo/tree/main/polling")
}
