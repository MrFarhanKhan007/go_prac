package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const EBTFile = "FileHandlingFiles\\EBTFile.txt"
const taxAmountFile = "FileHandlingFiles\\taxAmountFile.txt"
const EATFile = "FileHandlingFiles\\EATFile.txt"
const RatioFile = "FileHandlingFiles\\RatioFile.txt"

func readFromFileProfit(fileName string) (amount float64, err error) {
	amountInBytes, err := os.ReadFile(fileName)
	if err != nil {
		errorMsg := fmt.Sprintf("Something went wrong in reading the file - '%v' !", fileName)
		return 0, errors.New(errorMsg)
	}

	amountInString := string(amountInBytes)

	amount, err = strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return 0, errors.New("Something went wrong in parsing the value!")
	}
	return amount, nil
}

func writeToFileProfit(fileName string, data float64) {
	dataInBytes := fmt.Sprint(data)
	os.WriteFile(fileName, []byte(dataInBytes), 0644)
}

func profitCalculatorAdvanced() {
	// Revenue, Expenses and TaxRate
	revenue := getUserInputProfit("Enter Revenue: ")
	expenses := getUserInputProfit("Enter Expenses: ")
	taxRate := getUserInputProfit("Enter Tax Rate: ")

	fmt.Print("==================================\n==================================\n")

	calcValuesAndWriteToFile(revenue, expenses, taxRate)

	EBT, err := readFromFileProfit(EBTFile)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
	taxAmount, err := readFromFileProfit(taxAmountFile)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
	EAT, err := readFromFileProfit(EATFile)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
	Ratio, err := readFromFileProfit(RatioFile)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	printFormattedTextProfit("Earnings before tax", EBT)
	printFormattedTextProfit("Tax Amount", taxAmount)
	printFormattedTextProfit("Earnings after tax", EAT)
	printFormattedTextProfit("Ratio", Ratio)
}

func printTextProfit(text string) {
	fmt.Printf("%v", text)
}

func getUserInputProfit(text string) (userInput float64) {
	userInput = getUserInputProfitWithRetries(text, 3)
	return userInput
}

func getUserInputProfitWithRetries(text string, retries int) (userInput float64) {
	if retries == 0 {
		fmt.Print("Sorry, the maximum number of retries have been exhausted! Please try again later.\n")
		os.Exit(1)
	}
	printTextProfit(text)
	fmt.Scan(&userInput)
	if userInput < 0 {
		if retries != 1 {
			err := errors.New("Sorry, the respective input must be greater than 0! Kindly enter a valid input.\n")
			fmt.Println("ERROR! - ", err)
		}
		getUserInputProfitWithRetries(text, retries-1)
	}
	return userInput
}

func printFormattedTextProfit(text string, number float64) {
	fmt.Printf(text+": "+"%.2f\n", number)
}

func calcValuesAndWriteToFile(revenue, expenses, taxRate float64) {
	//Earnings before tax
	EBT := revenue - expenses
	writeToFileProfit(EBTFile, EBT)

	//Earnings after tax (profit)
	taxAmount := EBT * taxRate
	writeToFileProfit(taxAmountFile, taxAmount)

	EAT := EBT - taxAmount
	writeToFileProfit(EATFile, EAT)

	//Ratio
	Ratio := EBT / EAT
	writeToFileProfit(RatioFile, Ratio)
}
