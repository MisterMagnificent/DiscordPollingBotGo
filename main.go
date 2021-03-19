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
	blockingChan := make(chan bool, 1)
	killchan := make(chan os.Signal)
	signal.Notify(killchan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case sig := <-killchan:
			fmt.Printf("Killing (cleanly): %s\n", sig)
			polling.Cleanup()
			blockingChan <- true
			fmt.Println("leaving soon")
		}
	}()

	fmt.Println("preparing to block")
	<-blockingChan
	fmt.Println("after blockingChan, time to leave")
	return
}
