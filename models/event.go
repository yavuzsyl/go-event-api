package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	//later: add it to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
