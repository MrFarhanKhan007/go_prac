package main

import (
	"fmt"
	"math"
)

func investAmountCalculator() {
	var investAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = investAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
