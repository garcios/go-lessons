package models

import (
	"database/sql"
	"errors"
	"github.com/oskiegarcia/event-booking/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, date_time, user_id) 
VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	events := make([]Event, 0)

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

func GetEventByID(eventID int64) (*Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = ?`

	var event Event

	// use QueryRow when we expect a single row
	err := db.DB.QueryRow(query, eventID).Scan(&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `UPDATE events
    SET name = ?, description = ?, location = ?, date_time = ?
    WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}
