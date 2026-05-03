package Advanced

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func StructPrac() {
	UserfirstName := getUserData("Please enter your first name: ")
	UserlastName := getUserData("Please enter your last name: ")
	Userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(UserfirstName, UserlastName, Userbirthdate, time.Now())

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
	appUser.printUserData()
	appUser.clearUserData()
	appUser.printUserData()
}

func (u *user) printUserData() {
	fmt.Printf("First Name: %v\nLast Name: %v\nBirthDate: %v\nCreated at: %v\n", u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUserData() {
	u.firstName = "N/A"
	u.lastName = "N/A"
	u.birthDate = "N/A"
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func newUser(firstName string, lastName string, birthDate string, createdAt time.Time) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid Data!, Enter Valid Data please.")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: createdAt,
	}, nil
}
