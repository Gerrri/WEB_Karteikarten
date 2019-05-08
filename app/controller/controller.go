package controller

import(
	"net/http"
	"html/template"
)

//no Login
func Ausgeloggt_start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/ausgeloggt_main.tmpl", "./template/beides_startseite.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Ausgeloggt_karteikasten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/ausgeloggt_main.tmpl", "./template/ausgeloggt_karteikasten.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Ausgeloggt_registrieren(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/ausgeloggt_main.tmpl", "./template/ausgeloggt_registrieren.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

//Login
func Eingeloggt_start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/beides_startseite.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteikasten(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_karteikasten.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_meineKarteien(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_meinekarteien.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_meinProfil(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_meinprofil.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteiErstellen_01(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_karteierstellen_01.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteiErstellen_02(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_karteierstellen_02.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_karteikastenAnsehen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_karteikastenansehem.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_lernen_01(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_lernen_01.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_lernen_02(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_lernen_02.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}

func Eingeloggt_profilLoeschen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/eingeloggt_main.tmpl", "./template/eingeloggt_profilloeschen.tmpl")
	t.ExecuteTemplate(w, "layout", "")
}
