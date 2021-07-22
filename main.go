package main

import (
	"fmt"

	"github.com/Death4SwordLearnDev/GoBOT/bot"
	"github.com/Death4SwordLearnDev/GoBOT/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Printf("Failed to marshal config: %s\n", err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
