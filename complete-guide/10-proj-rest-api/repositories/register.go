package repositories

import (
	"database/sql"
	"errors"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"github.com/oskiegarcia/event-booking/db"
)

type Registration struct {
	ID      int64 `json:"id"`
	EventID int64 `json:"event_id"`
	UserID  int64 `json:"user_id"`
}

func (r *Registration) Save() error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.UserID, r.EventID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	r.ID = id

	return nil
}

func (r *Registration) Delete() error {
	query := `DELETE FROM registrations WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.ID)
	if err != nil {
		return err
	}

	return nil
}

func FindByUserIDAndEventID(userID, eventID int64) (*Registration, error) {
	query := `SELECT id FROM registrations WHERE user_id = ? AND event_id = ?`

	row := db.DB.QueryRow(query, userID, eventID)

	var registration Registration

	err := row.Scan(&registration.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.RegistrationNotFoundError
		}

		return nil, err
	}

	return &registration, nil
}
