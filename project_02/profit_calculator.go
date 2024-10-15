package main

import (
	"fmt"
	"os"
)

func writeResultsToFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.2f\nPROFIT: %.2f\nRATIO: %.2f", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(result), 0644)
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	
	revenue = getValueFromConsole("Provide revenue: ")
	expenses = getValueFromConsole("Provide expenses: ")
	taxRate = getValueFromConsole("Provide tax rate: ")

	earningsBeforeTax, profit, ratio := calcAllValues(
		revenue,
		expenses,
		taxRate,
	)

	writeResultsToFile(earningsBeforeTax, profit, ratio)

	outPutTextAndValue("Earnings before tax: ", earningsBeforeTax)
	outPutTextAndValue("Earnings after tax: ", profit)
	outPutTextAndValue("Ratio: ", ratio)
}

func getValueFromConsole(prompt string) float64 {
	currentValue := 0.0

	for {
		fmt.Print(prompt)
		fmt.Scan(&currentValue)

		if currentValue > 0.0 {
			return currentValue
		}

		fmt.Println("Invalid input, try again!")
	}

}

func calcAllValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func outPutTextAndValue(prompt string, value float64)  {
	formattedValue := fmt.Sprintf("%v: %.2f", prompt, value)

	fmt.Println(formattedValue)
}
