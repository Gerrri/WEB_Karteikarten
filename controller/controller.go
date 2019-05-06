package controller

import(
	"net/http"
	"html/template"
)

/* ######################   not logged in Pages   ###################### */
func Ausgeloggt_start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/ausgeloggt_main.html", "./Templates/beides_startseite.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Ausgeloggt_karteikasten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/ausgeloggt_main.html", "./Templates/ausgeloggt_karteikasten.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Ausgeloggt_registrieren(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/ausgeloggt_main.html", "./Templates/ausgeloggt_registrieren.html")
	t.ExecuteTemplate(w, "layout", "")
}

/* ######################   logged in Pages   ###################### */
func Eingeloggt_start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/beides_startseite.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteikasten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_karteikasten.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_meineKarteien(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_meinekarteien.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_meinProfil(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_meinprofil.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteiErstellen_01(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_karteierstellen_01.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteiErstellen_02(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_karteierstellen_02.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteikastenAnsehen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_karteikastenansehem.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_lernen_01(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_lernen_01.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_lernen_02(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggt_lernen_02.html")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_profilLoeschen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/eingeloggt_main.html", "./Templates/eingeloggtprofilloeschen.html")
	t.ExecuteTemplate(w, "layout", "")
}
