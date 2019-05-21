package controller

import (
	"fmt"
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

type tmp_L_modkarteikasten1 struct {
	Karteien              string
	AlleKarten            []Karte
	AlleFortschirtte      []int
	AktuelleKarte         Karte
	AktuellerKarteikasten Karteikasten
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
	t.ExecuteTemplate(w, "layout", nil)
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
	t.ExecuteTemplate(w, "layout", nil)
}

func L_lernen(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_lernen.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func L_meinekarteikaesten(w http.ResponseWriter, r *http.Request) {

	data := tmp_L_MeineKarteikaesten{
		Karteien:                  strconv.Itoa(GetKarteikastenAnz()),
		GespeicherteKarteikaesten: []Karteikasten{},
		MeineKarteikaesten:        []Karteikasten{},
	}

	nutzer := GetNutzerById(1) //muss noch dynamisch gehlot werden

	for _, element := range nutzer.ErstellteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(1)))
		data.MeineKarteikaesten = append(data.MeineKarteikaesten, temp_kk)
	}

	for _, element := range nutzer.GelernteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(1)))
		data.GespeicherteKarteikaesten = append(data.GespeicherteKarteikaesten, temp_kk)
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten.html")
	t.ExecuteTemplate(w, "layout", data)
}

func L_meinProfil(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func L_meinProfil_popup(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil_popup.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func L_modkarteikasten1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten1.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func L_modkarteikasten2(w http.ResponseWriter, r *http.Request) {
	data := tmp_L_modkarteikasten1{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		AktuellerKarteikasten: Karteikasten{},
		AlleKarten:            []Karte{},
	}

	temp_kk := GetKarteikastenByid(1)
	temp_kk.FortschrittP = int(GetKarteikastenFortschritt(GetKarteikastenByid(1), GetNutzerById(1)))

	data.AktuellerKarteikasten = temp_kk

	for i, element := range temp_kk.Karten {
		data.AlleKarten = append(data.AlleKarten, element)
		data.AlleKarten[i].Num = i + 1
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten2.html")
	t.ExecuteTemplate(w, "layout", data)
}

func L_showKarteikarten(w http.ResponseWriter, r *http.Request) {

	data := tmp_L_modkarteikasten1{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		AktuellerKarteikasten: Karteikasten{},
		AlleKarten:            []Karte{},
		AlleFortschirtte:      []int{},
		AktuelleKarte:         Karte{},
	}

	temp_kk := GetKarteikastenByid(1)
	temp_kk.FortschrittP = int(GetKarteikastenFortschritt(GetKarteikastenByid(1), GetNutzerById(1)))

	//gewählte Karte

	Num := r.FormValue("Num")
	if Num == "" {
		Num = "1"
	}

	data.AktuellerKarteikasten = temp_kk

	for i, element := range temp_kk.Karten {
		data.AlleKarten = append(data.AlleKarten, element)
		data.AlleKarten[i].Num = i + 1
	}

	for _, element := range GetKarteikastenWiederholungArr(temp_kk, GetNutzerById(1)) {
		data.AlleFortschirtte = append(data.AlleFortschirtte, element)
	}

	akt, _ := strconv.Atoi(Num)
	akt = akt - 1

	//fmt.Println("#########################################################################################")
	//fmt.Println(akt)
	//fmt.Println("#########################################################################################")
	data.AktuelleKarte = data.AlleKarten[akt]
	data.AktuelleKarte.NutzerFach = strconv.Itoa(data.AlleFortschirtte[akt])

	fmt.Println(data.AktuelleKarte.NutzerFach)

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_showKarteikarten.html")
	t.ExecuteTemplate(w, "layout", data)
}
