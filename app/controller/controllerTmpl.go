package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var SessionNutzerID = ""
//var EinlogName = 

type tmp_b_home struct {
	Nutzer     string
	Lernkarten string
	Karteien   string
}

type tmp_L_lernen struct {
	//Menüleiste
	Nutzer     string
	Lernkarten string
	Karteien   string

	//Kasten
	Name           string
	Kategorie      string
	UnterKategorie string
	Fortschritt    int
	Kartenwd       [5]int
	Kartenanz      int

	//Karte
	Titel   string
	Frage   string
	Antwort string

	//nächste Karte
	KartenID     int
	KastenID     string
	NextKartenID int
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
	KastenID                  string
	KartenID                  int
}

type tmp_L_modkarteikasten1 struct {
	Karteien              string
	AlleKarten            []Karte
	Wiederholungen        []int
	AktuelleKarte         Karte
	AktuellerKarteikasten Karteikasten
	//von aktueller Karte
	Titel   string
	Frage   string
	Antwort string

	//aktueller Kasten
	KastenID string
	KartenID int
}

/* ######################   not logged in Pages   ###################### */
func NL_Home(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST" {

		r.ParseForm()
		var nutzername=r.FormValue("nutzername")
		var passwort=r.FormValue("passwort")
		var nutzer=GetAlleNutzer()

		for _, arr:=range nutzer{
			if arr.Name==nutzername && arr.Name!=""{
				if arr.Passwort==passwort{
					SessionNutzerID=arr.DocID
					fmt.Println(SessionNutzerID)
					r.Method=""
					L_meinekarteikaesten(w, r)
				}else{
					p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
					t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in.html")

					t.ExecuteTemplate(w, "layout", p)
				}
				
			}else{
				p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
				t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in.html")

				t.ExecuteTemplate(w, "layout", p)
			}
		
		}
		
	}else{
		p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
		t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in.html")

		t.ExecuteTemplate(w, "layout", p)
	}
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
	kk = GetAlleKarteikaestenOeffentlich()

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
<<<<<<< HEAD
=======
	if r.Method == "POST" {
		r.ParseForm()
		var benutzername = r.FormValue("benutzername")
		var email = r.FormValue("email")
		var passwort = r.FormValue("passwort")
		var passwortWdhl = r.FormValue("passwortWdhl")
		var datenschutz = r.FormValue("datenschutz")

		fmt.Println(datenschutz)

		var nutzer = GetAlleNutzer()

		var vorhanden_mail bool = false
		var vorhanden_benutzer bool = false
		var passwortUngleich bool = false

		for _, arr := range nutzer {
			if arr.EMail == email {
				vorhanden_mail = true
			}
			if arr.Name == benutzername {
				vorhanden_benutzer = true
			}
		}

		if passwort != passwortWdhl {
			passwortUngleich = true
		}

		fmt.Println("E-Mail vorhanden: ", vorhanden_mail)
		fmt.Println("Passwoerter nicht gleich: ", passwortUngleich)

		if !vorhanden_mail && !vorhanden_benutzer && !passwortUngleich && datenschutz == "on" {
			var hinzufuegen Nutzer
			hinzufuegen.EMail = email
			hinzufuegen.Name = benutzername
			hinzufuegen.Passwort = passwort
			hinzufuegen.TYP = "nutzer"
			hinzufuegen.ErstellteKarteien = []string{}
			hinzufuegen.GelernteKarteien = []string{}
			AddNutzer(hinzufuegen)
			NL_Home(w, r)
		}

	}
>>>>>>> master
	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_registrieren.html")

	t.ExecuteTemplate(w, "layout", p)
}

/* ######################   logged in Pages   ###################### */
func L_Home(w http.ResponseWriter, r *http.Request) {
	
	p := tmp_b_home{Nutzername: GetNutzerById(SessionNutzerID), Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/b_home.html", "./templates/L_logged_in.html")

	t.ExecuteTemplate(w, "layout", p)

}

func L_karteikaesten(w http.ResponseWriter, r *http.Request) {
	data := tmp_nL_Karteikasten{
		Nutzername: GetNutzerById(SessionNutzerID),
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		Naturwissenschaften:   []Karteikasten{},
		Sprachen:              []Karteikasten{},
		Gesellschaft:          []Karteikasten{},
		Wirtschaft:            []Karteikasten{},
		Geisteswissenschaften: []Karteikasten{},
		Sonstige:              []Karteikasten{},
	}

	kk := []Karteikasten{}
	kk = GetAlleKarteikaestenOeffentlich()

	//fmt.Println("kk[]:", kk)

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
	var query = r.URL.Query()

	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])
	var kasten = GetKarteikastenByid(Kastenid)
	var karte = kasten.Karten[Kartenid]

	data := tmp_L_lernen{
		//Allgemein
		Nutzername: GetNutzerById(SessionNutzerID),
		Nutzer:     strconv.Itoa(GetNutzeranz()),
		Lernkarten: strconv.Itoa(GetKartenAnz()),
		Karteien:   strconv.Itoa(GetKarteikastenAnz()),

		//Kasten
		Name:           kasten.Titel,
		Kategorie:      kasten.Kategorie,
		UnterKategorie: kasten.Unterkategorie,
		Fortschritt:    kasten.FortschrittP,
		Kartenwd:       [5]int{0, 0, 0, 0, 0},
		Kartenanz:      len(kasten.Karten),

		//Karte
		Titel:   karte.Titel,
		Frage:   karte.Frage,
		Antwort: karte.Antwort,

		//nächste karte
		KartenID: Kartenid + 1,
		KastenID: Kastenid,
	}

	//fmt.Printf("KartenID: %v\n", data.KartenID)
	//fmt.Printf("größe: %v\n", (len(kasten.Karten) - 1))

	if data.KartenID >= (len(kasten.Karten)) {
		data.KartenID = 0
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_aufdecken.html")

	t.ExecuteTemplate(w, "layout", data)
}

func L_lernen(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	//Kastenid und Kartenid auslesen
	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])

	var kasten = GetKarteikastenByid(Kastenid)
	var karte = kasten.Karten[Kartenid]

	//Ergebnis auswerten
	var erg = query["Ergebnis"]

	if erg != nil {
		var Ergebnis, err = strconv.Atoi(erg[0])
		if err == nil {
			//Richtig
			if Ergebnis == 1 {
				//Update Lernstatus
				fmt.Println("Richtig")

				UpdateKarteikastenKarte(Kastenid, Kartenid-1, GetNutzerById(SessionNutzerID), true)

				//fmt.Println("KastenID: %v", Kastenid)
				//fmt.Println("KartenID: %v", Kartenid-1)
				//fmt.Println("NutzerID: %v", 1)

			}

			//Falsch
			if Ergebnis == 2 {
				//Update Lernstatus
				fmt.Println("Falsch")

				UpdateKarteikastenKarte(Kastenid, Kartenid-1, GetNutzerById(SessionNutzerID), false)
				//fmt.Println("KastenID: ", Kastenid)
				//fmt.Println("KartenID: ", Kartenid-1)
				//fmt.Println("NutzerID: ", 1)
			}
		}
	}
	//Data für Template
	data := tmp_L_lernen{
		//Allgemein
		Nutzername: GetNutzerById(SessionNutzerID),
		Nutzer:     strconv.Itoa(GetNutzeranz()),
		Lernkarten: strconv.Itoa(GetKartenAnz()),
		Karteien:   strconv.Itoa(GetKarteikastenAnz()),

		//Kasten
		Name:           kasten.Titel,
		Kategorie:      kasten.Kategorie,
		UnterKategorie: kasten.Unterkategorie,
		Fortschritt:    kasten.FortschrittP,
		Kartenwd:       [5]int{0, 0, 0, 0, 0},
		Kartenanz:      len(kasten.Karten),

		//Karte
		Titel:   karte.Titel,
		Frage:   karte.Frage,
		Antwort: karte.Antwort,

		//nächste karte
		KartenID:     Kartenid,
		KastenID:     Kastenid,
		NextKartenID: Kartenid + 1,
	}

	if data.KartenID >= (len(kasten.Karten)) {
		data.KartenID = 0
	}

	if data.NextKartenID >= (len(kasten.Karten)) {
		data.NextKartenID = 0
	}

	//fmt.Printf("%vHier: \n", data.Titel)

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_lernen.html")

	t.ExecuteTemplate(w, "layout", data)
}

func L_meinekarteikaesten(w http.ResponseWriter, r *http.Request) {

	data := tmp_L_MeineKarteikaesten{
		Karteien:                  strconv.Itoa(GetKarteikastenAnz()),
		GespeicherteKarteikaesten: []Karteikasten{},
		MeineKarteikaesten:        []Karteikasten{},
	}

	nutzer := GetNutzerById(SessionNutzerID) //muss noch dynamisch gehlot werden

	titel := ""
	beschreibung := ""
	kategorie := ""
	radio := ""
	if r.Method == "POST" {

		r.ParseForm()
		titel = r.FormValue("titel")
		fmt.Println(titel)
		beschreibung = r.FormValue("beschreibung")
		fmt.Println(beschreibung)
		kategorie = r.FormValue("kategorie")
		fmt.Println(kategorie)
		radio = r.FormValue("answer")
		fmt.Println(radio)

		OberKategorie := ""

		if kategorie == "Biologie" || kategorie == "Chemie" || kategorie == "Elektrotechnik" || kategorie == "Informatik" || kategorie == "Mathematik" || kategorie == "Medizin" || kategorie == "Naturkunde" || kategorie == "Physik" {
			OberKategorie = "Naturwissenschaften"
		}
		if kategorie == "Chinesisch" || kategorie == "Deutsch" || kategorie == "Englisch" || kategorie == "Französisch" || kategorie == "Griechisch" || kategorie == "Italienisch" || kategorie == "Latein" || kategorie == "Russisch" {
			OberKategorie = "Sprachen"
		}
		if kategorie == "Ethik" || kategorie == "Geschichte" || kategorie == "Literatur" || kategorie == "Musik" || kategorie == "Politik" || kategorie == "Recht" || kategorie == "Soziales" || kategorie == "Sport" || kategorie == "Verkehrskunde" {
			OberKategorie = "Gesellschaft"
		}
		if kategorie == "BWL" || kategorie == "Finanzen" || kategorie == "Landwirtschaft" || kategorie == "Marketing" || kategorie == "VWL" {
			OberKategorie = "Wirtschaft"
		}
		if kategorie == "Kriminologie" || kategorie == "Philosophie" || kategorie == "Psychologie" || kategorie == "Pädagogik" || kategorie == "Theologie" {
			OberKategorie = "Geisteswissenschaften"
		}
		if kategorie == "Sonstige" {
			OberKategorie = "Sonstige"
		}

		kk := Karteikasten{}
		kk.TYP = "Karteikasten"
		kk.NutzerID = nutzer.DocID
		kk.Sichtbarkeit = radio
		kk.Kategorie = OberKategorie
		kk.Unterkategorie = kategorie
		kk.Titel = titel
		kk.Anzahl = 0
		kk.Beschreibung = beschreibung

		AddKarteikasten(kk, nutzer)
		//FUNKTIONIERT noch nicht
	}

	for _, element := range nutzer.ErstellteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(SessionNutzerID)))
		data.MeineKarteikaesten = append(data.MeineKarteikaesten, temp_kk)

	}

	for _, element := range nutzer.GelernteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(SessionNutzerID)))
		data.GespeicherteKarteikaesten = append(data.GespeicherteKarteikaesten, temp_kk)

	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten.html")
	t.ExecuteTemplate(w, "layout", data)
}

func L_meinProfil(w http.ResponseWriter, r *http.Request) {
	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil.html")

	t.ExecuteTemplate(w, "layout", p)
}


func L_meinProfil_popup(w http.ResponseWriter, r *http.Request) {
	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil_popup.html")

	t.ExecuteTemplate(w, "layout", p)
}

func L_modkarteikasten1(w http.ResponseWriter, r *http.Request) {
	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten1.html")

	t.ExecuteTemplate(w, "layout", p)
}

// NutzerID?
func L_modkarteikasten2(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])

	//Post auswertung
	if r.Method == "POST" {

		r.ParseForm()
		typ := r.FormValue("type")
		fmt.Println("type: ", typ)

		r.ParseForm()
		titel := r.FormValue("titel")
		//fmt.Println(titel)

		r.ParseForm()
		frage := r.FormValue("frage")
		//fmt.Println(frage)

		r.ParseForm()
		antwort := r.FormValue("antwort")
		//fmt.Println(antwort)

		//Save option bei "+"

		if typ == "mod" {
			fmt.Println("Update:", titel, frage, antwort)
			UpdateKarteikarte(Kastenid, Kartenid, titel, frage, antwort)
		}

		if typ == "add" {
			fmt.Println("Add:", titel, frage, antwort)
			AddKarteikarte(Kastenid, titel, frage, antwort)
		}

		if typ == "del" {
			fmt.Println("Delete:", titel, frage, antwort)
			DelKarteikarteByID(Kastenid, Kartenid)

			Kartenid = 0
		}

	}

	var kasten = GetKarteikastenByid(Kastenid)
	var karte = kasten.Karten[Kartenid]

	data := tmp_L_modkarteikasten1{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		AktuellerKarteikasten: Karteikasten{},
		AlleKarten:            []Karte{},

		Titel:   karte.Titel,
		Frage:   karte.Frage,
		Antwort: karte.Antwort,

		KastenID: Kastenid,
		KartenID: Kartenid,
	}

	temp_kk := GetKarteikastenByid(Kastenid)
	temp_kk.FortschrittP = int(GetKarteikastenFortschritt(GetKarteikastenByid(Kastenid), GetNutzerById(SessionNutzerID)))

	data.AktuellerKarteikasten = temp_kk

	for i, element := range temp_kk.Karten {
		data.AlleKarten = append(data.AlleKarten, element)
		data.AlleKarten[i].Num = i + 1
		data.AlleKarten[i].Index = i
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten2.html")
	t.ExecuteTemplate(w, "layout", data)
}

//NutzerID austauschen
func L_showKarteikarten(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	//Kastenid und Kartenid auslesen
	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])

	var kasten = GetKarteikastenByid(Kastenid)
	var karte = kasten.Karten[Kartenid]

	data := tmp_L_modkarteikasten1{
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		AktuellerKarteikasten: Karteikasten{},
		AlleKarten:            []Karte{},
		Wiederholungen:        []int{},
		AktuelleKarte:         Karte{},
		KastenID:              Kastenid,
		KartenID:              Kartenid,
	}

	kasten.FortschrittP = int(GetKarteikastenFortschritt(kasten, GetNutzerById(SessionNutzerID)))

	//gewählte Karte

	Num := r.FormValue("Num")
	if Num == "" {
		Num = "1"
	}

	data.AktuellerKarteikasten = kasten

	for i, element := range kasten.Karten {
		data.AlleKarten = append(data.AlleKarten, element)
		data.AlleKarten[i].Index = i
		data.AlleKarten[i].Num = i + 1
	}

	//fmt.Println("kk: ", kasten)

	for _, element := range GetKKWiederholungenByNutzer(kasten, GetNutzerById(SessionNutzerID)) {
		data.Wiederholungen = append(data.Wiederholungen, element)
	}

	akt, _ := strconv.Atoi(Num)
	akt = akt - 1

	//fmt.Println("AlleFortschritte: ", data.Wiederholungen)
	//fmt.Println("akt: ", akt)

	data.AktuelleKarte = karte
	data.AktuelleKarte.NutzerFach = strconv.Itoa(data.Wiederholungen[akt])

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_showKarteikarten.html")
	t.ExecuteTemplate(w, "layout", data)
}
