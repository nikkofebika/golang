package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func JSONEncode(value any) {
	byte, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	JSONEncode("Nikko")
	JSONEncode(true)
	JSONEncode(1)
	JSONEncode([]string{"Nikko", "Golang"})
}

type Address struct {
	Street   string
	City     string
	PostCode int
}

type User struct {
	Name      string
	Age       int
	IsAdmin   bool
	Hobbies   []string
	Addresses []Address
}

func TestEncodeObject(t *testing.T) {
	var user User = User{
		Name:    "Nikko",
		Age:     20,
		IsAdmin: true,
		Hobbies: []string{"Golang", "PHP", "Javascript"},
		Addresses: []Address{
			{
				Street:   "Jalan Raya",
				City:     "Jakarta",
				PostCode: 12345,
			},
			{
				Street:   "Jalan Menuju Surga",
				City:     "Semarang",
				PostCode: 112233,
			},
		},
	}

	byte, _ := json.Marshal(user)

	fmt.Println(string(byte))
}
