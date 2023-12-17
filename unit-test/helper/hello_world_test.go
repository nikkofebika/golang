package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.M) {
	fmt.Println("Before Test")

	t.Run() // run all test

	fmt.Println("After Test")
}

func TestHelloWorld1(t *testing.T) {
	// if runtime.GOOS == "windows" {
	// 	t.Skip("can not run in windows")
	// }
	name := "Nikko"
	name2 := "Nikko Fe"
	result := HelloWorld(name)

	// assert.Equal(t, "Hello "+name2, result, "Result is not Hello "+name2)
	require.Equal(t, "Hello "+name2, result, "Result is not Hello "+name2)

	// if result != "Hello "+name2 {
	// 	// panic("Result is not Hello " + name2) // don't use panic
	// 	t.Fatal("Result is not Hello " + name2)
	// }

	fmt.Println("DONE TestHelloWorld1")
}

func TestHelloWorld2(t *testing.T) {
	// name := "Nikko"
	name2 := "Nikko Fe"
	result := HelloWorld(name2)

	if result != "Hello "+name2 {
		// panic("Result is not Hello " + name2)
		t.Fatal("Result is not Hello " + name2)
	}

	fmt.Println("DONE TestHelloWorld2")
}
