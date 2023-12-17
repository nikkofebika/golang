package main

import (
	"fmt"
	"net/http"
)

func main() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello Goleng</h1>")
	}
	// fmt.Println(reflect.TypeOf(handler))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
