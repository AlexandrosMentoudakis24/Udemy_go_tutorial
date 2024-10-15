package main

import "fmt"

type stringFloatMap map[string]float64

func (m stringFloatMap) Print() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	myCourses := []string{ "Angular", "Flutter", "React" }

	userNames = append(userNames, "Alexandros")
	userNames = append(userNames, "Lefteris")

	fmt.Println(userNames)
	fmt.Println(len(userNames))

	courseRatings := make(stringFloatMap, 5)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5

	courseRatings.Print()
	fmt.Println()

	for k, v := range myCourses {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	for k, v := range courseRatings {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

}

