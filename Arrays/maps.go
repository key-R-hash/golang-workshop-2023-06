package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon aws services": "https://www.aws.com",
	}

	fmt.Println(websites["google"])

	websites["linkdin"] = "https://linkdin.com"

	fmt.Println(websites["linkdin"])

	delete(websites, "linkdin")
	fmt.Println(websites)

}
