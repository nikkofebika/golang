package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	sm := http.NewServeMux()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home Page")
	})

	sm.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Page")
	})

	sm.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		fmt.Println(len(paths))
		fmt.Println(paths[len(paths)-1])
		fmt.Println(paths[0])
		fmt.Println(paths[1])
		fmt.Println(paths[2])
		fmt.Fprint(w, "Hello "+paths[len(paths)-1])
	})

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: sm,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
