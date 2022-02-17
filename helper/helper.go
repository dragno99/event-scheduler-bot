package helper

import (
	"errors"
	"strings"

	"google.golang.org/api/calendar/v3"
)

func GetEvent(data []string) (*calendar.Event, error) {

	if len(data) != 4 {
		return nil, errors.New("enter a valid input")
	}

	attendee := strings.SplitAfter(data[3], " ")

	event := &calendar.Event{
		Summary: strings.TrimSpace(data[0]),
		Start: &calendar.EventDateTime{
			DateTime: strings.TrimSpace(data[1]) + "+05:30",
			TimeZone: "IST",
		},
		End: &calendar.EventDateTime{
			DateTime: strings.TrimSpace(data[2]) + "+05:30",
			TimeZone: "IST",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
		Attendees:  []*calendar.EventAttendee{},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: "<RANDOM_VALUE>",
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
				Status: &calendar.ConferenceRequestStatus{
					StatusCode: "success",
				},
			},
		},
	}
	for _, x := range attendee {
		event.Attendees = append(event.Attendees, &calendar.EventAttendee{
			Email: strings.TrimSpace(x),
		})
	}

	return event, nil
}
