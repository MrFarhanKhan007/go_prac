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

type Admin struct{
	email string
	password string
	User
}

func (u *User) PrintUserData() {
	fmt.Printf("First Name: %v\nLast Name: %v\nbirthDate: %v\nCreated at: %v\n", u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserData() {
	u.firstName = "N/A"
	u.lastName = "N/A"
	u.birthDate = "N/A"
}

func NewUser(firstName string, lastName string, birthDate string, createdAt time.Time) (*User, error) {
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

func NewAdmin(email string, password string)(*Admin, error){
	if email == "" || password == ""  {
		return nil, errors.New("Invalid Data!, Enter Valid Data please.")
	}
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "DAY-ONE",
			createdAt: time.Now(),
		},
	},nil
}