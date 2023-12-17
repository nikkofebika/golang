package main

import "fmt"

func loop(number int) int {
	total := 1
	for i := number; i > 0; i-- {
		total *= i
	}

	return total
}

func recrusiveLoop(number int) int {
	if number == 1 {
		return 1
	}
	return number * recrusiveLoop(number-1)
}

func main() {
	fmt.Println(loop(10))
	fmt.Println(recrusiveLoop(10))
}
