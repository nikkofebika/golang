package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "nikko"
	names[1] = "fe"

	fmt.Println(names)
	fmt.Println(names[1])

	// var values [3]int = [3]int{
	var values = [3]int{
		1,
		2,
	}
	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(len(values))

	values2 := [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(values2)
	fmt.Println(values2[1])
	fmt.Println(len(values2))

	// var slice []int = []int{
	slice := []int{
		1, 2, 3,
	}
	fmt.Println("slice")
	fmt.Println(slice)
	fmt.Println(slice[1])
	fmt.Println(len(slice))
}
