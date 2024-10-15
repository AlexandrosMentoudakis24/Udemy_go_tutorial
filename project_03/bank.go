package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"example/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("----------", err)
	}

	fmt.Println("Welcome to MyBank!")
	fmt.Println("Contact us on:", randomdata.PhoneNumber())

	for {

		displayMenu()

		userChoice := getUserChoice()

		if userChoice < 0 || userChoice > 4 {
			fmt.Println("Invalid userChoice, please select form 1 to 4!")

			continue
		}

		switch userChoice {
		case 1:
			fmt.Println("Account balance:", accountBalance)
		case 2:

			var depositAmount float64

			fmt.Print("Select amount to deposit: ")
			var _, _ = fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid ammount!")

				continue
			}

			accountBalance += depositAmount

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("Account balance:", accountBalance)
		case 3:
			var withdrawAmount float64

			fmt.Print("Select amount to withdraw: ")
			var _, _ = fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid ammount!")

				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient account balance!")

				continue
			}

			accountBalance -= withdrawAmount

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("Account balance:", accountBalance)
		default:
			fmt.Println("Exiting...")
			return
		}
	}
}

func getUserChoice() int {
	var userChoice int

	fmt.Print("Your userChoice: ")

	fmt.Scan(&userChoice)

	return userChoice
}

