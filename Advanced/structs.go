package Advanced

import (
	"fmt"
	"github.com/MrFarhanKhan007/go_prac/Advanced/user"
	"time"
)

func StructPrac() {
	UserfirstName := getUserData("Please enter your first name: ")
	UserlastName := getUserData("Please enter your last name: ")
	Userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(UserfirstName, UserlastName, Userbirthdate, time.Now())

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	adminUser, err := user.NewAdmin("daddy@gmail.com", "UniversalDad")
	adminUser.PrintUserData()
	adminUser.ClearUserData()
	adminUser.PrintUserData()
	fmt.Println("--------------------------")
	appUser.PrintUserData()
	appUser.ClearUserData()
	appUser.PrintUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
