package models

import (
	"fmt"
	"time"

	db "example.com/gin-go-api/sql-db"
)

type Event struct {
	ID          int64
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	UserId      int64
}

var events = []Event{}

func (e Event) SaveToQL() error {
	// use waitgroup
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	cmd, err := db.SQL_DB.Prepare(query)
	if err != nil {
		return err
	}
	defer cmd.Close()
	result, err := cmd.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id

	events = append(events, e)
	fmt.Println(events)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELEC*FROM events"
	rows, err := db.SQL_DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return event, nil
}
