package lists

import "fmt"

type Product struct {
	id int
	title string
	price float64
}

// join two arrays

// func main() {
// 	prices := []int{ 1, 2, 3, 4, 5 }
// 	newPrices := []int{ 6, 7, 8, 9, 10 }
//
// 	prices = append(prices, newPrices...)
// 	fmt.Println(prices)
// }

func main() {
	myHobbies := [3]string{ "Football", "Coding", "Running" }

	// print the whole list and then the first element
	printData(myHobbies)
	printData(myHobbies[0])

	// print the second and third element from the array
	// in an other stand alone array

	slicedHobbies := myHobbies[1:]
	slicedHobbies = myHobbies[1:3]

	printData(slicedHobbies)

	// create another sliced list that contains
	// the first and the second element of
	// the initial array
	anotherSlicedHobbies1:= myHobbies[:2]
	anotherSlicedHobbies2 := myHobbies[0:2]
	printData(anotherSlicedHobbies1)
	printData(anotherSlicedHobbies2)

	anotherSlicedHobbies1 = anotherSlicedHobbies1[1:3]
	printData(anotherSlicedHobbies1)

	myDynamicGoals := []string{ "Learn Go", "Complete Course" }
	printData(myDynamicGoals)

	myDynamicGoals[1] = "Pass all tests"
	printData(myDynamicGoals)

	myDynamicGoals = append(myDynamicGoals, "Random goal")
	printData(myDynamicGoals)

	myProducts := []Product{ 
		{ 1, "First Product", 25.25 },
		{ 2, "Second Product", 50.50 },
		{ 3, "Third Product", 10.10 },
	}
	printData(myProducts)

	newProduct := Product{ 4, "Fourth Product", 75.75 }

	myProducts = append(myProducts, newProduct)
	printData(myProducts)
}

func printData(data any) {
	fmt.Println(data)
}

