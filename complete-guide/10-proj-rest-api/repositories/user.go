package repositories

import (
	"database/sql"
	"errors"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"github.com/oskiegarcia/event-booking/db"
	"github.com/oskiegarcia/event-booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID

	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.InvalidCredentialsError
		}

		return err
	}

	if !utils.CheckPasswordHash(u.Password, hashedPassword) {
		return apperrors.InvalidCredentialsError
	}

	return nil
}
