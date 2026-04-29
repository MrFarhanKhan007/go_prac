package main

import "fmt"

func bankApplication() {
	accountBalance := 1000.0
	mainMenu()
	choice := getChoice()

	if choice == 1 {
		checkBalance(accountBalance)
	} else if choice == 2 {
		depositMoney(accountBalance)
	} else if choice == 3 {
		// TODO()
	} else {
		return
	}
}

func mainMenu() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Where do you want to go?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit ")
	fmt.Println("=============================")
}

func getChoice() int {
	var choice int
	fmt.Println("Your Choice: ")
	fmt.Scan(&choice)
	return choice
}

func checkBalance(accountBalance float64) {
	fmt.Println("Your Balance is: ", accountBalance)
}
func depositMoney(accountBalance float64) {
	var moneyToBeDeposited float64
	fmt.Scan(&moneyToBeDeposited)
	accountBalance += moneyToBeDeposited
}
