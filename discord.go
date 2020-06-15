package main

import (
	"fmt"
	"./config"
	"./bot"
)

func main() {
	err := config.ReadConfig()
	HandleError(err)

	bot.Start()

	<-make(chan struct{})
	return
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
