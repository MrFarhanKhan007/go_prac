package main

import (
	"fmt"
)

func profitCalculatorAdvanced() {
	// revenue, expenses and tax rate(in decimal)

	var revenue, expenses, taxRate float64

	printTextProfit("Enter Revenue: ")
	scanTextProfit(&revenue)

	printTextProfit("Enter Expenses: ")
	scanTextProfit(&expenses)

	printTextProfit("Enter Tax Rate: ")
	scanTextProfit(&taxRate)

	fmt.Println("==================================")
	fmt.Println("==================================")

	//Earnings before tax
	EBT := revenue - expenses
	printFormattedTextProfit("Earnings before tax", EBT)

	//Earnings after tax (profit)
	taxAmount := EBT * taxRate
	printFormattedTextProfit("Tax Amount", taxAmount)

	EAT := EBT - taxAmount
	printFormattedTextProfit("Earnings after tax", EAT)

	//Ratio
	ratio := EBT / EAT
	printFormattedTextProfit("Ratio", ratio)
}

func printTextProfit(text string) {
	fmt.Print(text)
}

func scanTextProfit(value *float64) {
	fmt.Scan(value)
}

func printFormattedTextProfit(text string, number float64) {
	fmt.Printf(text+": "+"%.2f\n", number)
}
