package controller

import (
	"html/template"
	"net/http"
	"strconv"
)

type tmp_b_home struct {
	Nutzer     string
	Lernkarten string
	Karteien   string
}

type tmp_nL_Karteikasten struct {
	Karteien              string
	Naturwissenschaften   []Karteikasten
	Sprachen              []Karteikasten
	Gesellschaft          []Karteikasten
	Wirtschaft            []Karteikasten
	Geisteswissenschaften []Karteikasten
	Sonstige              []Karteikasten
}

type tmp_L_MeineKarteikaesten struct {
	Karteien                  string
	GespeicherteKarteikaesten []Karteikasten
	MeineKarteikaesten        []Karteikasten
}

/* ######################   not logged in Pages   ###################### */
func NL_Home(w http.ResponseWriter, r *http.Request) {

	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in.html")

	t.ExecuteTemplate(w, "layout", p)
}

func NL_karteikaesten(w http.ResponseWriter, r *http.Request) {
	data := tmp_nL_Karteikasten{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		Naturwissenschaften:   []Karteikasten{},
		Sprachen:              []Karteikasten{},
		Gesellschaft:          []Karteikasten{},
		Wirtschaft:            []Karteikasten{},
		Geisteswissenschaften: []Karteikasten{},
		Sonstige:              []Karteikasten{},
	}

	kk := []Karteikasten{}
	kk = GetAlleKarteikaesten()

	for _, element := range kk {
		if element.Kategorie == "Naturwissenschaften" {
			data.Naturwissenschaften = append(data.Naturwissenschaften, element)
		} else if element.Kategorie == "Sprachen" {
			data.Sprachen = append(data.Sprachen, element)
		} else if element.Kategorie == "Gesellschaft" {
			data.Gesellschaft = append(data.Gesellschaft, element)
		} else if element.Kategorie == "Wirtschaft" {
			data.Wirtschaft = append(data.Wirtschaft, element)
		} else if element.Kategorie == "Geisteswissenschaften" {
			data.Geisteswissenschaften = append(data.Geisteswissenschaften, element)
		} else {
			data.Sonstige = append(data.Sonstige, element)
		}
	}

	t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_karteikaesten.html")
	t.ExecuteTemplate(w, "layout", data)
}

func NL_registrieren(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_registrieren.html")
	t.ExecuteTemplate(w, "layout", "")
}

/* ######################   logged in Pages   ###################### */
func L_Home(w http.ResponseWriter, r *http.Request) {
	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/b_home.html", "./templates/L_logged_in.html")

	t.ExecuteTemplate(w, "layout", p)
}

func L_karteikaesten(w http.ResponseWriter, r *http.Request) {
	data := tmp_nL_Karteikasten{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		Naturwissenschaften:   []Karteikasten{},
		Sprachen:              []Karteikasten{},
		Gesellschaft:          []Karteikasten{},
		Wirtschaft:            []Karteikasten{},
		Geisteswissenschaften: []Karteikasten{},
		Sonstige:              []Karteikasten{},
	}

	kk := []Karteikasten{}
	kk = GetAlleKarteikaesten()

	for _, element := range kk {
		if element.Kategorie == "Naturwissenschaften" {
			data.Naturwissenschaften = append(data.Naturwissenschaften, element)
		} else if element.Kategorie == "Sprachen" {
			data.Sprachen = append(data.Sprachen, element)
		} else if element.Kategorie == "Gesellschaft" {
			data.Gesellschaft = append(data.Gesellschaft, element)
		} else if element.Kategorie == "Wirtschaft" {
			data.Wirtschaft = append(data.Wirtschaft, element)
		} else if element.Kategorie == "Geisteswissenschaften" {
			data.Geisteswissenschaften = append(data.Geisteswissenschaften, element)
		} else {
			data.Sonstige = append(data.Sonstige, element)
		}
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_karteikaesten.html")
	t.ExecuteTemplate(w, "layout", data)
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

	data := tmp_L_MeineKarteikaesten{
		Karteien:                  strconv.Itoa(GetKarteikastenAnz()),
		GespeicherteKarteikaesten: []Karteikasten{},
		MeineKarteikaesten:        []Karteikasten{},
	}

	nutzer := GetNutzerById(1) //muss noch dynamisch gehlot werden

	for _, element := range nutzer.ErstellteKarteien {
		data.MeineKarteikaesten = append(data.MeineKarteikaesten, GetKarteikastenByid(element))
	}

	for _, element := range nutzer.GelernteKarteien {
		data.GespeicherteKarteikaesten = append(data.GespeicherteKarteikaesten, GetKarteikastenByid(element))
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten.html")
	t.ExecuteTemplate(w, "layout", data)
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
