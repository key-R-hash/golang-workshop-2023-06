package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("Hello World", "", "green", true)

	myFigure.Print()
}
