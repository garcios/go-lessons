package main

import (
	"fmt"

	"github.com/oskiegarcia/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating new user:", err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	fmt.Println("-----------------")
	appUser.OutputUserDetails()

	admin, err := user.NewAdmin("oscar@gmail.com", "pass123")
	if err != nil {
		fmt.Println("Error creating new admin:", err)
	}

	fmt.Println("-----------------")
	fmt.Printf("admin first name: %s\n", admin.GetFirstName())
	admin.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
