package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)

	if err == nil {
		fmt.Println("hasil ", hasil)
	} else {
		fmt.Println("error bro ", err.Error())
	}
}
