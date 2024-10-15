package user

import (
	"fmt"
	"errors"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	UserData
}

func (user *UserData) OutputUserData() {
	fmt.Println("First-Name: ", user.firstName)
	fmt.Println("Last-Name: ", user.lastName)
	fmt.Println("Birth-Date: ", user.birthDate)
	fmt.Println("Created-At: ", user.createdAt)
}

func (user *UserData) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		UserData: UserData{
			firstName: "Alex",
			lastName: "Mentou",
			birthDate: "24/01/2000",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (*UserData, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid inputs!")
	}

	createdAt := time.Now()


	return &UserData {
		firstName,
		lastName,
		birthDate,
		createdAt,
	}, nil
}
