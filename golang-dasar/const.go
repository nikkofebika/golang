package main

import "fmt"

func main() {
	// const firstName string = "nikko"
	// const lastName = "fe"
	// const age int = 63

	const (
		firstName string = "nikko"
		lastName         = "fe"
		age       int    = 63
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}
