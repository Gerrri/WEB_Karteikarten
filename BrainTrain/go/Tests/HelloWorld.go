package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler_H)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handler_H(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
