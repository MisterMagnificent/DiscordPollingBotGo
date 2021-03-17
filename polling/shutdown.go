package polling

import (
	"github.com/bwmarrin/discordgo"
	"syscall"
)

func shutdown(session *discordgo.Session, pollByChannel map[string]Poll) {
	for channel, poll := range pollByChannel {
		unpin(poll, session)
		//Write poll to memory here to pick it back up on set up
		_, _ = session.ChannelMessageSend(channel, "https://64.media.tumblr.com/d23f48a7699883b01bea3799a9d7c165/tumblr_pjdvccM2pA1rs5nuyo7_250.gif")
	}
}

func forceShutdown(session *discordgo.Session, pollByChannel map[string]Poll) {
	shutdown(session, pollByChannel)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
