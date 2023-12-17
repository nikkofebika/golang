package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	registerUser("nikko", func(name string) bool {
		return name == "asu"
	})

	filter := func(name string) bool {
		return name == "asu"
	}
	registerUser("asu", filter)
}
