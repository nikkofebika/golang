package main

import "fmt"

func sayHello(firstName string, lastName string) string {
	// func sayHello(firstName string, lastName string) {
	// fmt.Println("Hello", firstName, lastName)
	return "Hello " + firstName + " " + lastName
}

func getFullName(firstName string, lastName string) (string, string) {
	// func getFullName() (string, string) {
	// return "Nikko", "Fe"
	return firstName, lastName
}

func getNamedReturnValues() (firstName, lastName string, age int) {
	firstName = "nikko"
	lastName = "fe"
	age = 100
	return firstName, lastName, age
}

func main() {
	// sayHello("nikko", "fe")
	// sayHello("aaa", "bbb")

	// result := sayHello("nikko", "fe")
	// var result string = sayHello("nikko", "fe")
	// fmt.Println(result)

	// firstName, lastName := getFullName("nikko", "fe")
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName("nikko", "fe")
	fmt.Println(firstName)

	a, b, c := getNamedReturnValues()
	fmt.Println(a, b, c)
}
