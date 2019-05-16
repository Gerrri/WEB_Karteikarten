package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type tmp_b_home struct {
	Nutzer     string
	Lernkarten string
	Karteien   string
}

/* ######################   not logged in Pages   ###################### */
func NL_Home(w http.ResponseWriter, r *http.Request) {
	p := tmp_b_home{Nutzer: "asdasd", Lernkarten: "some news", Karteien: "Hannes"}
	t, err := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/b_home.html")
	fmt.Println(err)
	t.ExecuteTemplate(w, "layout", p)
}

func NL_karteikaesten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_karteikaesten.html")
	t.ExecuteTemplate(w, "layout", "")
}

func NL_registrieren(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_registrieren.html")
	t.ExecuteTemplate(w, "layout", "")
}

/* ######################   logged in Pages   ###################### */
func L_Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/b_home.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_karteikaesten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_karteikaesten.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_aufdecken(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_aufdecken.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_lernen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_lernen.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_meinekarteikaesten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_meinProfil(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_meinProfil_popup(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil_popup.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_modkarteikasten1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten1.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_modkarteikasten2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten2.html")
	t.ExecuteTemplate(w, "layout", "")
}

func L_showKarteikarten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_showKarteikarten.html")
	t.ExecuteTemplate(w, "layout", "")
}
