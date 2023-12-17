package main

import (
	"fmt"
)

func main() {
	days := [...]string{"monday", "tuesday", "wednesday", "thrusday", "friday", "saturday", "sunday"}
	// slice := []string{
	var slice []string = []string{
		"nikko", "fe",
	}
	fmt.Println(days)
	fmt.Println(len(days))
	fmt.Println(slice)

	// sliceDay1 := days[0:3]
	var sliceDay1 []string = days[5:]
	sliceDay1[0] = "saturday day"
	fmt.Println("days", days)
	fmt.Println("sliceDay1", sliceDay1)
	fmt.Println("sliceDay1 len", len(sliceDay1))
	fmt.Println("sliceDay1 cap", cap(sliceDay1))

	sliceDay2 := append(sliceDay1, "new day")
	fmt.Println("sliceDay2", sliceDay2)
	sliceDay2[0] = "sabtu"
	fmt.Println("sliceDay2", sliceDay2)
	fmt.Println("sliceDay2 len", len(sliceDay2))
	fmt.Println("sliceDay2 cap", cap(sliceDay2))
	fmt.Println("days", days)
	fmt.Println("sliceDay1", sliceDay1)

	var makeSlice = make([]string, 3, 5)
	makeSlice[0] = "a"
	makeSlice[1] = "b"
	makeSlice[2] = "c"
	fmt.Println("makeSlice", makeSlice)
	fmt.Println("makeSlice len", len(makeSlice))
	fmt.Println("makeSlice cap", cap(makeSlice))

	makeSlice2 := append(makeSlice, "d")
	fmt.Println("makeSlice2", makeSlice2)
	fmt.Println("makeSlice2 len", len(makeSlice2))
	fmt.Println("makeSlice2 cap", cap(makeSlice2))
	makeSlice2[0] = "x"
	fmt.Println("makeSlice2", makeSlice2)
	fmt.Println("makeSlice", makeSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("fromSlice", fromSlice)
	fmt.Println("toSlice", toSlice)
}
