package main

import (
	"fmt"
	"math"
)

func investAmountCalculatorUserInput() {
	const inflationRate = 2.5
	var investAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	realValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(realValue)
}
