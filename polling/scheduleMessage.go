package polling

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

// Schedule a message, format is UTC time.  D:HH:MM; where day is 0-6 i.e. Sunday-Saturday
func scheduleMessage(session *discordgo.Session, channelID string, options string, author string) {
	if author == "" {
		var optionList = strings.Split(options, " ")
		var option = optionList[0]
		var function = ""
		if len(optionList) > 1 {
			function = optionList[1]
		}

		if option != "" {
			var split = strings.Split(option, ":")

			if len(split) > 2 {
				weekday, err := strconv.Atoi(split[0])
				if err == nil && weekday < 7 {
					hour, err := strconv.Atoi(split[1])
					if err == nil && hour < 24 {
						minute, err := strconv.Atoi(split[2])
						if err == nil && minute < 60 {
							fmt.Println("Scheduled: %s", split)

							switch function {
							case "result":
								(*ourScheduler).Every().Weekday(weekday).Hour(hour).Minute(minute).Second(0).Do(getResultSched, pollByChannel[channelID], session)
							default:
								function = "view"
								(*ourScheduler).Every().Weekday(weekday).Hour(hour).Minute(minute).Second(0).Do(view, pollByChannel[channelID], session)
							}
							_, _ = session.ChannelMessageSend(channelID, "A "+function+" is scheduled for weekday: "+split[0]+", hour: "+split[1]+" (UTC), and minute: "+split[2]+"; once a week")
						}
					}
				}
			}
		}
	} else {
		var delayList = strings.Split(options, " ")
		var delayDays, err = strconv.Atoi(delayList[0])
		var delayHours = 0
		if err == nil && len(delayList) > 1 {
			delayHours, err = strconv.Atoi(delayList[1])
		}

		var remindFunc = func(channelID string, author string) {
			allowArray := [](discordgo.AllowedMentionType){}
			allowArray = append(allowArray, discordgo.AllowedMentionTypeUsers)
			var messageSend = &discordgo.MessageSend{
				Content: "You're reminded, <@" + author + ">!",
				AllowedMentions: &discordgo.MessageAllowedMentions{
					Parse: allowArray,
				},
			}
			_, _ = session.ChannelMessageSendComplex(channelID, messageSend)
		}
		if delayDays > 0 {
			(*ourScheduler).Delay().Day(delayDays).Hour(delayHours).Do(remindFunc, channelID, author)
		} else {
			(*ourScheduler).Delay().Hour(delayHours).Do(remindFunc, channelID, author)
		}
	}
}
