package models

import (
	"errors"
	"eventapi/db"
	"time"
)

type Event struct {
	ID          int64
	Title       string
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) SetId(id int64) {
	e.ID = id
}

func (e *Event) Save() error {
	//later: add it to database
	command := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(command)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.SetId(id)
	return err
}

func (e Event) Update() error {
	command := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id= ? 
	`
	stmt, err := db.DB.Prepare(command)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&e.Name, &e.Description, &e.Location, &e.DateTime, &e.ID)

	return err
}

func (e Event) Delete() error {
	deleteCommand := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(deleteCommand)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Register(userId int64) error {
	insertCommand := "INSERT INTO registrations(user_id, event_id) VALUES(?, ?)"

	stmnt, err := db.DB.Prepare(insertCommand)
	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(userId, e.ID)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	deleteCommand := "DELETE FROM registrations where user_id = ? and event_id = ?"
	stmnt, err := db.DB.Prepare(deleteCommand)
	if err != nil {
		return err
	}

	defer stmnt.Close()

	res, err := stmnt.Exec(userId, e.ID)
	if err != nil {
		return err
	}

	if aff, err := res.RowsAffected(); err != nil || aff <= 0 {
		return errors.New("error occured")
	}

	return nil
}
