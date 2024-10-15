package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":			   "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)

	websites["LinkedIn"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "LinkedIn")

	fmt.Println(websites)

	for k, v := range websites {
		fmt.Printf("Key: %v => Value: %v\n", k, v)
	}
}

