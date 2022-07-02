package bot

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"github.com/dragno99/event-scheduler-bot/config"
	"github.com/dragno99/event-scheduler-bot/scheduler"
)

var BotID string
var goBot *discordgo.Session

var mutex sync.Mutex

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
	go HandleRequest(s, m)
}

func HandleRequest(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.SplitAfter(m.Content, "\n")
	command := strings.TrimSpace(content[0])
	content = append(content[0:0], content[1:]...)

	switch command {
	case "!hey":
		{
			PrintRequest(s, m, "Hey "+m.Author.Username)
		}
	case "!schedule":
		{
			PrintRequest(s, m, scheduler.ScheduleEvent(content))
		}
	case "!upcoming":
		{
			PrintRequest(s, m, scheduler.ShowUpcomingEvent())
		}
	case "!update":
		{

			if len(content) != 2 {
				PrintRequest(s, m, "Enter a valid input.")
				return
			}

			idx, err := strconv.Atoi(strings.TrimSpace(content[0]))

			if err != nil || idx > len(scheduler.UpcomingEvents) || idx < 1 {
				PrintRequest(s, m, "Enter a valid input.")
				return
			}

			content[1] = strings.TrimSpace(content[1])
			PrintRequest(s, m, scheduler.AddAttendeesInEvent(scheduler.UpcomingEvents[idx-1].EventId, content[1]))

		}
	case "!delete":
		{
			if len(content) != 1 {
				PrintRequest(s, m, "Enter a valid input.")
				return
			}

			idx, err := strconv.Atoi(strings.TrimSpace(content[0]))

			if err != nil || idx > len(scheduler.UpcomingEvents) || idx < 1 {
				PrintRequest(s, m, "Enter a valid input.")
				return
			}
			PrintRequest(s, m, scheduler.DeleteEvent(scheduler.UpcomingEvents[idx-1].EventId))

		}
	}

}

func PrintRequest(s *discordgo.Session, m *discordgo.MessageCreate, message string) {

	mutex.Lock()
	_, _ = s.ChannelMessageSend(m.ChannelID, message)
	mutex.Unlock()

}
