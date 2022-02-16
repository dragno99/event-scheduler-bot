package main

import (
	"fmt"

	"github.com/dragno99/event-scheduler-bot/bot"
	"github.com/dragno99/event-scheduler-bot/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
