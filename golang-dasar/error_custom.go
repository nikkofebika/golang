package main

import (
	"fmt"
	"reflect"
)

type errorValidation struct {
	Message string
}

func (err *errorValidation) Error() string {
	return err.Message
}

type notFoundValidation struct {
	Message string
}

func (err notFoundValidation) Error() string {
	return err.Message
}

func saveData(name string, data any) error {
	if name == "" {
		return &errorValidation{"error validation"}
	}

	if name != "nikko" {
		return notFoundValidation{"error not found"}
	}
	return nil
}

func main() {
	// errValidation := errorValidation{"message ni bos"}
	// fmt.Println(errValidation)
	// fmt.Println(errValidation.Message)
	// fmt.Println(errValidation.Error())

	err := saveData("", nil)
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err))

	if err != nil {
		// switch errorResult := err.(type) {
		// case *errorValidation:
		// 	fmt.Println("error errorValidation ", errorResult.Error())
		// case *notFoundValidation:
		// 	fmt.Println("error notFoundValidation ", errorResult.Error())
		// default:
		// 	fmt.Println("Unknown error")
		// }

		if errorResult, ok := err.(*errorValidation); ok {
			fmt.Println("error errorValidation ", errorResult.Error())
		} else if errorRes, ok := err.(notFoundValidation); ok {
			fmt.Println("error notFoundValidation ", errorRes.Error())
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("save data success")
	}

}
