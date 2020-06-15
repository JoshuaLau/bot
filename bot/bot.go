package bot

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotID string
var bot *discordgo.Session

func Start() {
	fmt.Println("Starting session..")
	bot, err := discordgo.New("Bot " + config.Token)
	HandleError(err)

	user, err := bot.User("@me")
	HandleError(err)

	BotID = user.ID

	bot.AddHandler(MessageHandler)
	err = bot.Open()
	HandleError(err)

	fmt.Printf("Bot %s is running!\n", user)

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !strings.HasPrefix(m.Content, config.BotPrefix) { 
		return
	}

	if m.Author.Bot {
		return
	}

	text := strings.Split(m.Content, " ")
	if len(text) == 1 {
		return
	}

	command := text[0][1:]

	argument := strings.Join(text[1:], " ")
	_, _ = s.ChannelMessageSend(m.Message.ChannelID, argument)

}

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
