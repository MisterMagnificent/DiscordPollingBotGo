package main

import (
	"./config"
	"./polling"
	"fmt"
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
