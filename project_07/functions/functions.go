package functions

import "fmt"

type transformFn func(int, int) int

func main() {
	numbers := []int{ 1, 2, 3, 4 }

	doubledNumbers := transformNumbers(&numbers, 5, transform)

	fmt.Println(numbers)
	fmt.Println(doubledNumbers)

}

func transformNumbers(
	numbers *[]int,
	multiplier int,
	transform transformFn,
) []int {
	dNumbers := make([]int, len(*numbers))

	for index, value := range *numbers {
		dNumbers[index] = transform(value, multiplier)
	}

	return dNumbers
}

func transform(value, multiplier int) int {
	return value * multiplier
}
