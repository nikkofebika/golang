package test

import (
	"fmt"
	"learn/api/simple"
	"testing"
)

func TestHelloManual(t *testing.T) {
	sayHello := simple.NewSayHello()
	helloService := simple.NewSayHelloService(sayHello)

	fmt.Println(sayHello)
	fmt.Println(helloService)
}

func TestHelloWire(t *testing.T) {
	helloService := simple.InitializeHelloService()

	// fmt.Println(sayHello)
	fmt.Println(helloService)
}
