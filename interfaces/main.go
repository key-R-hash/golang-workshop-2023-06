package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

type logUser struct {
	firstname string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed")
	}
}

func (user logUser) log() {
	fmt.Println(user.firstname)
}

func main() {
	userLog := &logData{"User data logged", "data-log.txt"}
	// do more work ...
	createLog(userLog)

	message := logUser{"User logged in"}
	// do more work ...
	createLog(message)

	outputValue(4)
	outputValue("kiarash")
}

func createLog(data logger) {
	// more things we do
	data.log()
}

func outputValue(value interface{}) {
	val, ok := value.(string)

	fmt.Println(val, ok)
}
