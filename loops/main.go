package main

import (
	"fmt"
	"time"
)

func main() {
	productName := [5]string{"book", "laptop", "headphone"}

	for i := 0; i < 5; i++ {
		fmt.Println(productName[i])
	}

	for index, value := range productName {
		fmt.Printf("Index: %v, Value: %v\n", index, value)

		if index == 2 {
			// continue
			break
		}

		fmt.Println(productName[index])
	}

	for true {
		time.Sleep(time.Second * 1)
		fmt.Println("infinite loop")
	}
}
