package main

import "fmt"

type Person struct {
	Name string
	Age  int16
}

func (person *Person) setName() {
	person.Name = "Mr. " + person.Name
	person.Age = 100
}

func main() {
	// person := Person{}
	// setName(&person)
	// var person *Person = &Person{}
	var person Person = Person{Name: "Nikko"}
	person.setName()

	fmt.Println(person)
}
