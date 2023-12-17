package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "1234567890"

	var cthKTP string = "0987654321"
	var ktp2 NoKTP = NoKTP(cthKTP)

	fmt.Println(ktp)
	fmt.Println(ktp2)

}
