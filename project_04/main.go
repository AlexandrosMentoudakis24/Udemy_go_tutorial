package main

import (
	"fmt"
	"example/pointers/user"
)

func main() {
	userFirstName := getUserData("Provide your first name: ")
	userLastName := getUserData("Provide your last name: ")
	userBirthDate := getUserData("Provide your birth date (DD/MM/YYYY): ")

	var user1 *user.UserData

	user1, error := user.New(userFirstName, userLastName, userBirthDate)

	if error != nil {
		panic("Failed to create new User!")
	}

	user1.OutputUserData()
	user1.ClearUserName()
	user1.OutputUserData()

	admin := user.NewAdmin("test@test.com", "12345")

	admin.OutputUserData()
	admin.ClearUserName()
	admin.OutputUserData()
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)

	return userInput
}
