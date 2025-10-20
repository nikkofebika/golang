package web

import (
	"fmt"
	"net/http"
	"testing"
)

// type Middleware func(next http.Handler) http.Handler

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BEFORE EXECUTING", r.URL.Path)
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("AFTER EXECUTING", r.Proto)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO WORLD KU")
	})

	mux.HandleFunc("/silite/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO SILITE")
	})

	mux.HandleFunc("/silite/rio/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO SILITERIO")
	})

	handlerMiddleware := LogMiddleware{Handler: mux}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &handlerMiddleware,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO WORLD KU")
	})

	mux.HandleFunc("/silite/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO SILITE")
	})

	mux.HandleFunc("/silite/rio/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO SILITERIO")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}

func TestHandler(t *testing.T) {
	HandFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HELLO WORLD KU")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: HandFunc,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}
