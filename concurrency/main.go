package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)

	go generateValue(channel)
	go generateValue(channel)

	sum := <-channel + <-channel

	fmt.Println(sum)
}

func generateValue(c chan int) {
	fmt.Println("Generating value...")
	sleepTime := rand.Intn(3) + 1

	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := rand.Intn(10)
	c <- result
}

// func main() {
// 	greet()
// 	storeData("This is some dummy data!", "dummy-data.txt")

// 	channel := make(chan int)

// 	go storeMoreData(50000, "50000_1.txt", channel)
// 	go storeMoreData(50000, "50000_2.txt", channel)

// 	<-channel
// 	<-channel
// }

// func greet() {
// 	fmt.Println("Hi there!")
// }

// func storeData(storableText string, fileName string) {
// 	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 	defer file.Close()

// 	if err != nil {
// 		fmt.Println("Creating the file failed. Exiting.", err)
// 		return
// 	}

// 	_, err = file.WriteString(storableText)

// 	if err != nil {
// 		fmt.Println("Writing to the file failed.")
// 	}
// }

// func storeMoreData(lines int, fileName string, c chan int) {
// 	for i := 0; i < lines; i++ {
// 		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
// 		storeData(text, fileName)
// 	}

// 	fmt.Printf("-- Done storing %v lines of text --\n", lines)

// 	c <- 1
// }
