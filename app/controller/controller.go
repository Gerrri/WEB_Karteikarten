package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = `<html>
<head>
    <title>Hello World!</title>
</head>
<body>
    {{ . }}
</body>
</html>
`

type template_x struct {
	Title string
	News  string
}

type TnL_home struct {
	Nutzer     int
	Lernkarten int
	Karteien   int
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func Test_site(w http.ResponseWriter, r *http.Request) {
	p := template_x{Title: "Eine temmplate Test Seite :)", News: "JAPPP"}
	t, _ := template.ParseFiles("./templates/template.html")
	t.Execute(w, p)
}

func NL_Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nL_Home")
	p := TnL_home{Nutzer: 123, Lernkarten: 312, Karteien: 27}
	t, _ := template.ParseFiles("./templates/TnL_home.html")
	fmt.Println(t.Execute(w, p))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	fmt.Fprint(w, "login")
}

func NL_karteikaesten(w http.ResponseWriter, r *http.Request) {
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	t.Execute(w, "My first Template ")
}