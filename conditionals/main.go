package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge, err := getUserAge()

	if err != nil {
		fmt.Println(err)
	}

	if (userAge > 20 && userAge < 30) || userAge == 40 {
		fmt.Println("Welcome to the club")
	} else if userAge > 50 {
		fmt.Println("The best age")
	} else {
		fmt.Println("Sorry, you are not old enough")
	}

	fmt.Println("Goodbye!")
}

func getUserAge() (int64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your age: ")

	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return userAge, err
}
