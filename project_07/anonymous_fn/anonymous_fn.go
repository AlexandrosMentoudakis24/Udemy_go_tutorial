package anonymous_fn

import "fmt"

func main() {
	numbers := []int{ 1, 2, 3 }

	double := createTransformer(2)
	triple := createTransformer(3)

	doubledUp := transformNumbers(
		&numbers,
		double,
	)

	tripledUp := transformNumbers(
		&numbers,
		triple,
	)

	fmt.Println(numbers)
	fmt.Println(doubledUp)
	fmt.Println(tripledUp)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(multiplier int) func(int) int {
	return func(number int) int {
		return number * multiplier
	}
}


