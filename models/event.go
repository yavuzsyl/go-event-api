package models

import "time"

type Event struct {
	ID          int64
	Title       string
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e Event) Save() {
	//later: add it to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
