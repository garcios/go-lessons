package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Admin Type embedding
type Admin struct {
	email    string
	password string
	User
}

// New is constructor function.
func New(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("invalid firstName or lastName")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email string, password string) (*Admin, error) {
	user, err := New("admin", "admin", "---")
	if err != nil {
		return nil, err
	}
	admin := &Admin{
		email:    email,
		password: password,
		User:     *user,
	}
	return admin, nil
}

func (u *User) OutputUserDetails() {
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birthday: %s\n", u.birthdate)
	fmt.Printf("Created At: %s\n", u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (a *Admin) GetFirstName() string {
	return a.firstName
}
