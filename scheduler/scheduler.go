package scheduler

import (
	"google.golang.org/api/calendar/v3"
)

func Schedule(event *calendar.Event) string {
	srv, err := GetNewService()

	if err != nil {
		return "Faild to create event"
	}

	event, err = srv.Events.Insert("primary", event).Do()

	if err != nil {
		return "Faild to create event"
	}
	return "Event created"
}
