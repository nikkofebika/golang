package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func JsonEncode(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		panic(err)
	}

	return string(data)
}

func TestEncode(t *testing.T) {
	fmt.Println(JsonEncode("Nikko"))
	fmt.Println(JsonEncode(123))
	fmt.Println(JsonEncode(true))
	fmt.Println(JsonEncode([]string{"apple", "banana", "cherry"}))
}
