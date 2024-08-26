package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 2.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Provide investment amount: ")
	var _, _ = fmt.Scan(&investmentAmount)
	fmt.Print("Provide expected return rate: ")
	var _, _ = fmt.Scan(&expectedReturnRate)
	fmt.Print("Provide period in years: ")
	var _, _ = fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflation / 100, years)

	fmt.Println("Future invest value:", futureValue)
	fmt.Println("Future invest value after inflation:", futureRealValue)
}
