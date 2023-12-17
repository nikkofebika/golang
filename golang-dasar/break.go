package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		if counter == 5{
			break
		}

		fmt.Println("counter ", counter)
		counter++
	}
	fmt.Println("DONE")

	for i := 0; i < 10; i++ {
		if i == 5{
			continue
		}

		fmt.Println("i ", i)
	}
	fmt.Println("DONE")
}
