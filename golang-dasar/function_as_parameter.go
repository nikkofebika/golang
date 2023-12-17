package main

import "fmt"

type Filter func(string) string

// func sayHellow(name string, filter func(string) string) string {
func sayHellow(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func filterWord(name string) string {
	if name == "asu" {
		return "..."
	}

	return name
}

func main() {
	fmt.Println(sayHellow("nikko", filterWord))
	fmt.Println(sayHellow("asu", filterWord))

	sh := sayHellow
	fw := filterWord
	fmt.Println(sh("filter", fw))
}
