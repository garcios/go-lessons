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

type RegisteredUser struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

func (r *Registration) Save() error {
	query := "INSERT INTO registrations (user_id, event_id) VALUES (?, ?)"

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
	query := `DELETE FROM registrations WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.UserID, r.EventID)
	if err != nil {
		return err
	}

	return nil
}

func FindRegistrationByUserIDAndEventID(userID, eventID int64) (*Registration, error) {
	query := `SELECT id FROM registrations WHERE user_id = ? AND event_id = ?`

	row := db.DB.QueryRow(query, userID, eventID)

	registration := Registration{
		UserID:  userID,
		EventID: eventID,
	}

	err := row.Scan(&registration.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.RegistrationNotFoundError
		}

		return nil, err
	}

	return &registration, nil
}

func GetRegisteredUsersForEventID(eventID int64) ([]*RegisteredUser, error) {
	query := `SELECT users.id, users.email FROM users JOIN registrations 
    ON users.id = registrations.user_id WHERE registrations.event_id = ?`

	rows, err := db.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	registeredUsers := make([]*RegisteredUser, 0)

	for rows.Next() {
		var user RegisteredUser
		err := rows.Scan(&user.UserID, &user.Email)
		if err != nil {
			return nil, err
		}

		registeredUsers = append(registeredUsers, &user)
	}

	if err = rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return registeredUsers, nil
		}

		return nil, err
	}

	return registeredUsers, nil
}
