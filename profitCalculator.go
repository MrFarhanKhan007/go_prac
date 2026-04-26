package main

import "fmt"

func packageCalculator() {
	// revenue, expenses and tax rate(in decimal)

	var revenue, expenses, taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Println("==================================")
	fmt.Println("==================================")

	//Earnings before tax
	EBT := revenue - expenses
	fmt.Println("Earnings before tax: ",EBT)

	//Earnings after tax (profit)
	taxAmount := EBT * taxRate
	fmt.Println("Tax Amount: ",taxAmount)

	EAT := EBT - taxAmount
	fmt.Println("Earnings after tax: ",EAT)

	//Ratio
	ratio := EBT / EAT
	fmt.Println("Ratio: ",ratio)

}
