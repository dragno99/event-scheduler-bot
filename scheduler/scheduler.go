package scheduler

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dragno99/event-scheduler-bot/helper"
	"google.golang.org/api/calendar/v3"
)

var (
	UpcomingEvents []struct {
		Summary string
		EventId string
	}
)

func ScheduleEvent(content []string) string {
	srv, err := GetNewService()

	if err != nil {
		fmt.Println("Failed to create event service")
		return "Failed to create event."
	}

	event, err := helper.GetEvent(content)

	if err != nil {
		return "Enter a valid input."
	}

	event, err = srv.Events.Insert("primary", event).ConferenceDataVersion(1).Do()

	if err != nil {
		fmt.Println(err.Error())
		return "Failed to create event."
	}
	fmt.Println("link", event.HangoutLink)

	return "Event successfully scheduled.\n\nGoogle meet's link: " + event.HangoutLink
}

func ShowUpcomingEvent() string {

	srv, err := GetNewService()

	if err != nil {
		return "Error occured."
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(50).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next 50 of the user's events: %v", err)
	}
	var res string
	res += "Upcoming events:\n\n"
	if len(events.Items) == 0 {
		return "No upcoming events found."
	} else {

		UpcomingEvents = append(UpcomingEvents[0:0], UpcomingEvents[len(UpcomingEvents):]...)

		for _, item := range events.Items {
			date := item.Start.DateTime
			UpcomingEvents = append(UpcomingEvents, struct {
				Summary string
				EventId string
			}{
				Summary: item.Summary,
				EventId: item.Id,
			})
			if date == "" {
				date = item.Start.Date
			}
			res += fmt.Sprintf("-\t%v (%v)\n\n", item.Summary, date)
		}
		return res
	}
}

func AddAttendeesInEvent(eventId string, attendee string) string {
	srv, err := GetNewService()

	if err != nil {
		fmt.Println(err.Error())
		return "Updation failed."
	}

	event, err := srv.Events.Get("primary", eventId).Do()

	if err != nil {
		fmt.Println(err.Error())
		return "Updation failed."
	}

	attendeeArray := strings.SplitAfter(attendee, " ")

	for _, x := range attendeeArray {
		event.Attendees = append(event.Attendees, &calendar.EventAttendee{
			Email: strings.TrimSpace(x),
		})
	}

	_, err = srv.Events.Update("primary", eventId, event).Do()

	if err != nil {
		fmt.Println(err.Error())
		return "Updation failed."
	}

	return "Event update successfully !!!"

}
