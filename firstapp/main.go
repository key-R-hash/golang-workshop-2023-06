package main

import "fmt"

func main() {
	// var GreetingText string
	// GreetingText = "Hello, world!"
	// var GreetingText string = "Hello, world!"
	// GreetingText := "Hello, world!"

	// luckeyNumber := 10
	// evenMoreLuckyNumber := luckeyNumber + 5
	// evenMoreLuckyNumber := luckeyNumber * 2

	// newNumber := luckeyNumber / 3
	// var newNumber float64 = float64(luckeyNumber) / 3
	// var newNumber float32 = float32(luckeyNumber) / 3

	// firstName := "kiarash"
	// lastName := "amiri"

	// fmt.Println(firstName + " " + lastName)

	// sum := add(4, 10)
	// fmt.Println(sum)

	// a, b := generateRandomNumber()
	// fmt.Println(b)
	// fmt.Println(a)
	// fmt.Println(a + b)

	// fmt.Println(greetingText)
	// fmt.Println(greeting.GreetingText)

	height, weight := getUserMetrics()
	bmi := weight / (height * weight)

	fmt.Println(bmi)
}
