package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dragno99/event-scheduler-bot/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content != "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
	} else if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
	}
}
