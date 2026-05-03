package Advanced

import (
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

	var appUser user = user{
		firstName: UserfirstName,
		lastName:  UserlastName,
		birthDate: Userbirthdate,
		createdAt: time.Now(),
	}
	printUserData(&appUser)
}

func printUserData(u *user) {
	fmt.Printf("First Name: %v\nLast Name: %v\nBirthDate: %v\nCreated at: %v", u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
