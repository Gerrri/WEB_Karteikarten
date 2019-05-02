package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type template_x struct {
	Title string
	News  string
}

func index(w http.ResponseWriter, r *http.Request) {
	p := template_x{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("./template.html")
	t.Execute(w, p)
}

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
