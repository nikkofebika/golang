package main

import (
	"fmt"
)

func main() {
	// var person map[string]string = map[string]string{
	// person := map[string]string{
	// 	"name":  "nikko",
	// 	"email": ".com",
	// 	"age":   "15",
	// }

	person := make(map[string]string)
	person["name"] = "nikko"
	person["email"] = ".com"
	person["address"] = "indo"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["email"])
	fmt.Println(person["address"])
}
