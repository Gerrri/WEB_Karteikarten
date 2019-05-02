package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {fmt.Fprint(w, "index")}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "login")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
