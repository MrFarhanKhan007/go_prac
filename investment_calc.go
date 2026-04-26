package main

import (
	"fmt"
	"math"
)

func investAmountCalculator() {
	// var investAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years = 10

	// var investAmount, expexpectedReturnRate float64 = 1000,5.5

	const inflationRate = 2.5
	investAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0
	futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	realValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(realValue)
}
