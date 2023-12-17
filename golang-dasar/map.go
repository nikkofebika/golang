package main

import "fmt"

func main() {
	// var map1 = map[string]string{
	// map1 := map[string]string{
	// 	"name":  "nikko",
	// 	"email": ".com",
	// 	"age":   "27",
	// }

	var map1 = make(map[string]string)
	map1["name"] = "nikko"
	map1["email"] = ".com"
	map1["age"] = "27"

	fmt.Println(map1)
	fmt.Println(len(map1))
	fmt.Println(map1["name"])

	delete(map1, "age")
	fmt.Println(len(map1))
	fmt.Println(map1)
}