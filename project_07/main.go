package main

import "fmt"

func main() {
	numbers := []int{ 1, 2, 3, 4, 5 }

	sum := sumup(1,2,3,4,5,6,7,8,9,)

	fmt.Println(numbers)
	fmt.Println(sum)

}

func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

