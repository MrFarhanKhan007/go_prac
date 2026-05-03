package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) PrintUserData() {
	fmt.Printf("First Name: %v\nLast Name: %v\nbirthDate: %v\nCreated at: %v\n", u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserData() {
	u.firstName = "N/A"
	u.lastName = "N/A"
	u.birthDate = "N/A"
}

func New(firstName string, lastName string, birthDate string, createdAt time.Time) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid Data!, Enter Valid Data please.")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: createdAt,
	}, nil
}
