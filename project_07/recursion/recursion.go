package recursion

import "fmt"

func main() {

	result := calcFuctorial(5)

	fmt.Println(result)
}

func calcFuctorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * calcFuctorial(number - 1)
}

