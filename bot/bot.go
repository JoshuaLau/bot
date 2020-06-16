package bot

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	BotID string
	bot *discordgo.Session
	commandToText map[string]string
)

func Start() {
	commandToText = map[string]string{
		"cc": "coding challenge",
		"interview": "interview",
		"offer": "offer",
		"rejection": "rejection",
		"reject": "rejection",
	}
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
		// add info about the command if its the only thing in the message
		return
	}

	command := text[0][1:]
	argument := strings.Join(text[1:], " ")
	formattedCommand := commandToText[command]

	if formattedCommand == "" {
		return
	}	
	if _, ok := config.Companies[argument]; !ok {
		s.ChannelMessageSend(m.Message.ChannelID, fmt.Sprintf("u idiot **%s** is not in the company list idiot!", argument))
		return
	}
	botText := fmt.Sprintf("Added **%s** for **%s**.", formattedCommand, argument)
	_, _ = s.ChannelMessageSend(m.Message.ChannelID, botText)

}

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
