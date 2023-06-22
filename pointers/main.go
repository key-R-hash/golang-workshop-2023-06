package main

import "fmt"

func main() {
	age := 40

	// fmt.Println(&age)

	// myage := &age
	var myage *int
	myage = &age

	*myage = 33

	// fmt.Println(*myage)

	fmt.Println(double(&age))
	fmt.Println(age)
}

func double(number *int) int {
	result := *number * 2
	*number = 100

	return result
}
