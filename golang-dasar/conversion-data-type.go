package main

import "fmt"

func main() {
	var int32 int32 = 299999999
	var int64 int64 = int64(int32)
	var int16 int16 = int16(int32)

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int16)

	var name = "nikko"
	var e = name[0]
	fmt.Print(name)
	fmt.Print(string(e))
	fmt.Print(e)
}
