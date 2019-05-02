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

type TnL_home struct {
	Nutzer     int
	Lernkarten int
	Karteien   int
}

func index(w http.ResponseWriter, r *http.Request) {
	p := template_x{Title: "asdAmazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("C:/Users/Dustin/Documents/GitHub/WEB_Karteikarten/go/Tests/template.html")
	t.Execute(w, p)
}

func nL_Home(w http.ResponseWriter, r *http.Request) {

	p := TnL_home{Nutzer: 22, Lernkarten: 312, Karteien: 27}
	t, _ := template.ParseFiles("C:/Users/Dustin/Documents/GitHub/WEB_Karteikarten/templates/TnL_home.html")
	t.Execute(w, p)

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	fmt.Fprint(w, "login")

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/nL_Home", nL_Home)

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()

}
