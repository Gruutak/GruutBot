package main

import (
	"gruutbot"
	"gruutbot/pkg/structs"
)

func main() {
	initConfig := structs.InitConfig{
		Prefix: "^",
	}

	bot := gruutbot.New(initConfig)
	bot.Start()
}
