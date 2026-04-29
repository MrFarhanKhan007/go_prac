package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func investAmountCalculatorAdvanced() {
	var investAmount, expectedReturnRate, years float64

	printTextInvestment("Enter investment amount: ")
	fmt.Scan(&investAmount)

	printTextInvestment("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	printTextInvestment("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, realValue := calcValueInvestment(investAmount, expectedReturnRate, years)

	printFormattedTextInvestment(futureValue)
	printFormattedTextInvestment(realValue)
}

func printFormattedTextInvestment(text float64) {
	fmt.Printf("%.2f\n", text)
}
func printTextInvestment(text string) {
	fmt.Print(text)
}

func calcValueInvestment(investAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
