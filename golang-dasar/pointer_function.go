package main

import "fmt"

type Person struct {
	Name string
	Age  int16
}

func setName(person *Person) {
	person.Name = "budi andux"
}

func main() {
	// person := Person{}
	// setName(&person)
	var person *Person = &Person{}
	setName(person)

	fmt.Println(person)
}
