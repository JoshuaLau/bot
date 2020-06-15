package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NzIxODA1NzA1MDY1MDA1MTI2.XuZ4Hg.5UY9cHVvVm1HQpcid2BVqMsH4a0"
var BotID string

func main() {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	user, err := discord.User("@me")

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	BotID = user.ID

	discord.AddHandler(MessageHandler)

	err = discord.Open()

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Printf("Bot %s is running!\n", user)

	<-make(chan struct{})

	return

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	if m.Content == "sigma" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "yeet")
	}

	fmt.Println(m.Content)
}
