package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func JsonDecode(value string, data any) {
	err := json.Unmarshal([]byte(value), data)
	if err != nil {
		panic(err)
	}
}

func TestDecode(t *testing.T) {
	jsonString := `{"Name":"Nikko","Age":20,"IsAdmin":true,"Hobbies":["Golang","PHP","Javascript"],"Addresses":[{"Street":"Jalan Raya","City":"Jakarta","PostCode":12345},{"Street":"Jalan Menuju Surga","City":"Semarang","PostCode":112233}]}`

	// user := new(User)
	user := &User{}
	fmt.Println(user)
	JsonDecode(jsonString, user)

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.IsAdmin)
	fmt.Println(user.Hobbies)
	fmt.Println(user.Addresses)

	for _, address := range user.Addresses {
		fmt.Printf("Street : %s, City : %s, PostCode : %d\n", address.Street, address.City, address.PostCode)
	}
}

func TestDecodeArray(t *testing.T) {
	jsonString := `[{"Street":"Jalan Raya","City":"Jakarta","PostCode":12345},{"Street":"Jalan Menuju Surga","City":"Semarang","PostCode":112233}]`

	addresses := new([]Address)
	fmt.Println(addresses)
	JsonDecode(jsonString, addresses)
	fmt.Println(addresses)
	fmt.Println((*addresses)[0])

	addresses2 := []Address{}
	fmt.Println(addresses2)
	JsonDecode(jsonString, &addresses2)
	fmt.Println(addresses2)
	fmt.Println(addresses2[1])
}
