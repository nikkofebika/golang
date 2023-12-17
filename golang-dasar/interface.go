package main

import "fmt"

type HasName interface {
	GetName() string
	// GetAddress() string
	// GetAge() int
}

func sayHello(value HasName) {
	fmt.Println("Hello ", value.GetName())
	// fmt.Println("Hello ", value)
}

// func sayAddress(value HasName) {
// 	fmt.Println("Address ", value.GetAddress())
// }
// func sayAge(value HasName) {
// 	fmt.Println("Age ", value.GetAge())
// }

type Person struct {
	Name, Address string
	Age           int
}

func (person Person) GetName() string {
	return person.Name
}

// func (person Person) GetAddress() string {
// 	return person.Address
// }

// func (person Person) GetAge() int {
// 	return person.Age
// }

func main() {
	person1 := Person{"Nikko", "Tangerang", 20}
	person2 := new(Person)
	fmt.Println(person1)
	fmt.Println(person2)
	// sayHello(person1)
	// sayAddress(person1)
	// sayAge(person1)
}
