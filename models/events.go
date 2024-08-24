package models

import (
	"fmt"
	"log"
	"time"

	db "example.com/gin-go-api/sql-db"
)

type Event struct {
	ID          int64
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e Event) SaveToQL() error {
	// use waitgroup
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	fmt.Println("SQL_DB vaal->>>>>>>>>>>", db.SQL_DB)
	cmd, err := db.SQL_DB.Prepare(query)
	if err != nil {
		fmt.Println("ERR while -PREPARING== to QL", err)
		return err
	}
	defer cmd.Close()

	// Generate the next user_id
	var maxUserID int64
	err = db.SQL_DB.QueryRow("SELECT IFNULL(MAX(user_id), 0) + 1 FROM events").Scan(&maxUserID)
	if err != nil {
		fmt.Println("Failed to generate user_id", err)
		return err
	}
	e.UserId = maxUserID

	e.DateTime = time.Now() // Just an example; in real cases, you'll parse the DateTime from the JSON

	result, err := cmd.Exec(e.Name, e.Description, e.Location, e.DateTime.Format(time.RFC3339), e.UserId)
	if err != nil {
		fmt.Println("ERR while --EXECUTING++== to QL", err)
		return err
	}
	id, err := result.LastInsertId()
	fmt.Println("=>, id", id)
	e.ID = id

	events = append(events, e)
	fmt.Println(events)
	return err
}

func GetAllQLEvents() ([]Event, error) {
	query := "SELECT*FROM events"
	fmt.Println("SQL_DB vaal->>>>>>>>>>>", db.SQL_DB)
	rows, err := db.SQL_DB.Query(query)
	if err != nil {
		fmt.Println("query err->", err)
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var tmpEvent Event
		var dateTimeStr string
		fmt.Println(rows)
		err := rows.Scan(&tmpEvent.ID, &tmpEvent.Name, &tmpEvent.Description, &tmpEvent.Location, &dateTimeStr, &tmpEvent.UserId)

		if err != nil {
			fmt.Println("Scan err->", err)
			return nil, err
		}
		tmpEvent.DateTime, err = time.Parse(time.RFC3339, dateTimeStr)
		if err != nil {
			log.Println("Date parsing error:", err)
		}
		events = append(events, tmpEvent)
	}
	fmt.Println(events)

	return events, nil
}

func GetQLEventsById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.SQL_DB.QueryRow(query, id)

	var event Event
	var dateTimeStr string
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)
	if err != nil {
		return nil, err
	}
	event.DateTime, err = time.Parse(time.RFC3339, dateTimeStr)

	return &event, nil

}

func (e Event) UpdateQLEvent(id int64) error {

	query := `UPDATE events 
	SET user_id=?, name=?, description=?, location=?, dateTime=?
	WHERE id = ?`

	stmt, err := db.SQL_DB.Prepare(query)

	if err != nil {
		fmt.Println("Prepare query ERR->", err)
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.UserId, e.Name, e.Description, e.Location, e.DateTime.Format(time.RFC3339), id)
	if err != nil {
		fmt.Printf("error while execution UPDATE query with %v with ERR->> %v", id, err)
		return err
	}
	fmt.Println(result)
	return nil
}

func (e Event) DeleteQLEvent(id int64) error {
	deleteEventSQL := `DELETE FROM events WHERE id = ?`
	result, err := db.SQL_DB.Exec(deleteEventSQL, id)
	if err != nil {
		fmt.Println("Err while delete from QL, ERR-->>", err)
		return err
	}
	fmt.Println(result)
	return nil

}

func handleError(stmt string, err error) error {
	if err != nil {
		fmt.Printf("%v err ocurred %v", stmt, err)
		return err
	}
	return nil
}
