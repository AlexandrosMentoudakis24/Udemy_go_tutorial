package main

import (
	"fmt"
	"math"
)

const inflation = 2.5

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outPutText("Provide investment amount: ")
	var _, _ = fmt.Scan(&investmentAmount)
	outPutText("Provide expected return rate: ")
	var _, _ = fmt.Scan(&expectedReturnRate)
	outPutText("Provide period in years: ")
	var _, _ = fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValue(
		investmentAmount,
		expectedReturnRate,
		years,
	)

	// first way of printing values
	// fmt.Println("Future invest value:", futureValue)
	// fmt.Println("Future invest value after inflation:", futureRealValue)

	// second way of printing values
	// fmt.Printf("Future Value: %.0f\n", futureValue)
	// fmt.Printf("Future Value (after inflation): %.1f\n", futureRealValue)

	// third way of printing values
	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (after inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outPutText(text string) {
	fmt.Print(text)
}

func calcFutureValue(invAmount, expReturnRate, years float64) (float64, float64) {
	fv := invAmount * math.Pow(1 + expReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflation / 100, years)
	
	return fv, rfv
}
