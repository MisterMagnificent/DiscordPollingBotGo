package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MisterMagnificient/DiscordPollingBotGo/config"
	"github.com/MisterMagnificient/DiscordPollingBotGo/polling"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	polling.Start()

	killchan := make(chan os.Signal)
	signal.Notify(killchan, syscall.SIGTERM)

	go func() {
		select {
		case sig := <-killchan:
			fmt.Printf("Killing (cleanly): %s", sig)
			polling.Cleanup()
		}
	}()

	<-make(chan struct{})
	return
}
