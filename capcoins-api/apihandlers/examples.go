package apihandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct model for HelloFunc example
type Hello struct {
	Message string
}

func ExampleEchoFunc(w http.ResponseWriter, r *http.Request) {

	message := "(echo) Default message"
	if len(r.URL.Query()["message"]) > 0 {
		message = r.URL.Query()["message"][0]
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}


func ExampleHelloFunc(w http.ResponseWriter, r *http.Request) {

	message := Hello{"(hello) Welcome to CapCoins API crafted with Golang!"}
	jsonMessage, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonMessage)
}


