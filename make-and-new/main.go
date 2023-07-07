package main

import "fmt"

func main() {

	// hobbies := []string{"Sports", "Reading"}
	// hobbies = append(hobbies, "Game")
	// hobbies := make([]string, 2, 10)
	// fmt.Println(hobbies)
	// fmt.Println(len(hobbies), cap(hobbies))

	// hobbies[0] = "Sports"

	// hobbies = append(hobbies, "Game")

	// fmt.Println(hobbies)

	number := new(int)
	fmt.Println(*number)

	anotherNumber := 0
	numberAddress := &anotherNumber
	fmt.Println(*numberAddress)

}
