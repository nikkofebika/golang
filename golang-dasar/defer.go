package main

import (
	"fmt"
	"strconv"
)

func logging() {
	fmt.Println("Logging App...")

	// message := recover()
	// if message != nil {
	if message := recover(); message != nil {
		fmt.Println("terjadi error ", message)
	}
}

func runApplication(err bool) {
	defer logging()
	fmt.Println("Application RUN")
	if err {
		panic("Error BOS")
	}
}

func main() {
	cond, _ := strconv.ParseBool("1")
	fmt.Println(cond)
	runApplication(cond)
}
