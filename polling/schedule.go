package polling

import (
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
			if err != nil {
				hour, err := strconv.Atoi(split[1])
				if err != nil {
					minute, err := strconv.Atoi(split[2])
					if err != nil {
						(*ourScheduler).Every().Weekday(weekday).Hour(hour).Minute(minute).Do(view, pollByChannel[channelID], session)
					}
				}
			}
		}
	}
}
