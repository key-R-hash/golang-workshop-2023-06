package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createAt  time.Time
}

func main() {
	firstName := getUserData("please enter your first name")
	lastName := getUserData("please enter your last name")
	birthday := getUserData("please enter your birthday")

	userdata := NewUser(firstName, lastName, birthday)

	userdata.outputUserData()
}

func getUserData(text string) string {
	fmt.Println(text)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}

func (userdata User) outputUserData() {
	fmt.Println(userdata.firstName, userdata.lastName, userdata.birthday, userdata.createAt)
}

func NewUser(firstName string, lastName string, birthday string) User {
	createAt := time.Now()

	return User{
		firstName,
		lastName,
		birthday,
		createAt,
	}
}
