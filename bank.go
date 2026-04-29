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
		withdrawMoney(accountBalance)
	} else if choice == 4 {
		fmt.Println("Thank you for using our application!")
	} else {
		fmt.Println("Not a Correct choice!, Kindly choose a valid option.")
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
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)
	return choice
}

func checkBalance(accountBalance float64) {
	fmt.Println("Your Balance is: ", accountBalance)
}
func depositMoney(accountBalance float64) {
	var moneyToBeDeposited float64
	fmt.Print("Enter the money to be deposited: ")
	fmt.Scan(&moneyToBeDeposited)
	accountBalance += moneyToBeDeposited
	fmt.Printf("Balance Updated! - New Account Balance is: %.2f", accountBalance)
}

func withdrawMoney(accountBalance float64) {
	withdrawMoneyWithRetries(accountBalance, 3)
}

func withdrawMoneyWithRetries(accountBalance float64, retries int) {
	if retries == 0 {
		fmt.Println("Too many failed attempts. Please try again later! ")
		return
	}

	var moneyToWithdraw float64
	fmt.Print("Enter the money to be withdrawn: ")
	fmt.Scan(&moneyToWithdraw)

	if moneyToWithdraw <= 0 {
		fmt.Println("Please enter an amount that's greater than 0!")
		withdrawMoneyWithRetries(accountBalance, retries-1)
	} else if moneyToWithdraw > accountBalance {
		fmt.Println("Insufficient funds! - Withdrawn money cannot be greater than account balance. Please try again!")
		withdrawMoneyWithRetries(accountBalance, retries-1)
	} else {
		accountBalance -= moneyToWithdraw
		fmt.Printf("Balance Updated! - New Account Balance is: %.2f", accountBalance)
	}
}
