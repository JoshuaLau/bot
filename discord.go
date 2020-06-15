package main

import (
	"fmt"
	"os"
	"github.com/bwmarrin/discordgo"
)

var token string
var BotID string

func main() {
	token = os.Getenv("BOT_TOKEN")

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
