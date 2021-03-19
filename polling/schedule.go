package polling

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func schedule(session *discordgo.Session, channelID string, content string) {
	index := strings.IndexByte(content, ' ')
	chars := []rune(content)
	option := string(chars[index+1:])

	if option != "" {
		var split = strings.Split(option, ":")

		if len(split) > 2 {
			weekday, err := strconv.Atoi(split[0])
			if err == nil {
				hour, err := strconv.Atoi(split[1])
				if err == nil {
					minute, err := strconv.Atoi(split[2])
					if err == nil {
						fmt.Println("Scheduled: %s", split)
						(*ourScheduler).Every().Weekday(weekday).Hour(hour).Minute(minute).Second(0).Do(view, pollByChannel[channelID], session)
						_, _ = session.ChannelMessageSend(channelID, "A view is scheduled for weekday: "+split[0]+", hour: "+split[1]+" (UTC), and minute: "+split[2]+"; once a week")
					}
				}
			}
		}
	}
}
