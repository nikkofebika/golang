package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

// struct method
func (person Person) sayHello(name string) string {
	fmt.Println("Hello " + name + ", My name is " + person.Name)
}

func main() {
	var person Person = Person{}

	fmt.Println(person)

	// person2 := Person{Name: "Nikko", Address: "Indo", Age: 20}
	person2 := Person{"Nikko", "Indo", 20}
	fmt.Println(person2)

	// person3 := Person{}
	var person3 Person
	person3.Name = "Name"
	person3.Address = "Address"
	person3.Age = 30
	fmt.Println(person3)

	person4 := Person{"Nikko"}
	person4.sayHello("Bud")
}
