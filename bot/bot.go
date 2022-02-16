package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dragno99/event-scheduler-bot/config"
	"github.com/dragno99/event-scheduler-bot/helper"
	"github.com/dragno99/event-scheduler-bot/scheduler"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("start", err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println("start2", err.Error())
		return
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println("here", err.Error())
		return
	}
	fmt.Println("Bot is running...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	content := strings.SplitAfter(m.Content, "\n")
	if content[0] != "!schedule" {
		return
	}
	content = append(content[0:0], content[1:]...)

	event, err := helper.GetEvent(content)

	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Enter a valid input")
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, scheduler.Schedule(event))

}
