package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	
	fmt.Print("Provide revenue: ")
	var _, _ = fmt.Scan(&revenue)

	fmt.Print("Provide expenses: ")
	var _, _ = fmt.Scan(&expenses)

	fmt.Print("Provide taxRate: ")
	var _, _ = fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := (revenue - expenses) * (taxRate / 100)

	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings before tax: ", earningsBeforeTax)
	fmt.Println("Earnings after tax: ", earningsAfterTax)
	fmt.Println("Ratio: ", ratio)
}
