package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("data ", counter)
		counter++
	}
	fmt.Println("DONE")

	for i := 0; i < 10; i++ {
		fmt.Println("for2 ", i)
	}
	fmt.Println("DONE")

	var person = map[string]string{
		"name":  "nikko",
		"email": ".com",
		"age":   "27",
	}
	for i, v := range person {
		fmt.Println(i, v)
	}

}
