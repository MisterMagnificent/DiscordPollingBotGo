package main

import (
	"fmt"

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

	<-make(chan struct{})
	return
}
