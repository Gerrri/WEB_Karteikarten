package controller

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
	fmt.Fprint(w, "index")
}

func test_site(w http.ResponseWriter, r *http.Request) {
	p := template_x{Title: "Eine temmplate Test Seite :)", News: "JAPPP"}
	t, _ := template.ParseFiles("./templates/template.html")
	t.Execute(w, p)
}

func nL_Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nL_Home")
	p := TnL_home{Nutzer: 22222, Lernkarten: 312, Karteien: 27}
	t, _ := template.ParseFiles("./templates/TnL_home.html")
	fmt.Println(t.Execute(w, p))

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	fmt.Fprint(w, "login")

}


