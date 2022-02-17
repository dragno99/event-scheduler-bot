package bot

import (
	"fmt"
	"strconv"
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

	content := strings.SplitAfter(m.Content, "\n")
	command := strings.TrimSpace(content[0])
	content = append(content[0:0], content[1:]...)

	if command == "!schedule" {

		event, err := helper.GetEvent(content)

		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Enter a valid input.")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, scheduler.ScheduleEvent(event))

	} else if command == "!upcoming" {
		_, _ = s.ChannelMessageSend(m.ChannelID, scheduler.ShowUpcomingEvent())
	} else if command == "!update" {

		if len(content) != 2 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Enter a valid input1.")
			return
		}

		idx, err := strconv.Atoi(strings.TrimSpace(content[0]))

		fmt.Println(idx)

		if err != nil || idx > len(scheduler.UpcomingEvents) || idx < 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Enter a valid input1.")
			return
		}

		content[1] = strings.TrimSpace(content[1])

		_, _ = s.ChannelMessageSend(m.ChannelID, scheduler.AddAttendeesInEvent(scheduler.UpcomingEvents[idx-1].EventId, content[1]))

	} else if command == "!hey" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hey "+m.Author.Username)
	}

}
