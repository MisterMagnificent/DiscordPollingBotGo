package polling

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"syscall"
)

func shutdown(session *discordgo.Session, pollByChannel map[string]Poll) {
	pollFile, _ := json.MarshalIndent(pollByChannel, "", " ")
	_ = ioutil.WriteFile("polls.json", pollFile, 0644)

	eventFile, _ := json.MarshalIndent(eventManagerByChannel, "", " ")
	_ = ioutil.WriteFile("events.json", eventFile, 0644)

	for channel, _ := range pollByChannel {
		_, _ = session.ChannelMessageSend(channel, "https://64.media.tumblr.com/d23f48a7699883b01bea3799a9d7c165/tumblr_pjdvccM2pA1rs5nuyo7_250.gif")
	}
	session.Close()
}

func forceShutdown(session *discordgo.Session, pollByChannel map[string]Poll) {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
