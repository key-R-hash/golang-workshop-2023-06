package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func add(num1 int, num2 int) int {
	result := num1 + num2

	return result
}

func generateRandomNumber() (int, int) {
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)

	return randomNumber1, randomNumber2
}

func getUserMetrics() (height float64, weight float64) {
	reader := bufio.NewReader(os.Stdin)

	weightInput, _ := reader.ReadString('\n')
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ = strconv.ParseFloat(weightInput, 64)
	height, _ = strconv.ParseFloat(heightInput, 64)

	return
}
