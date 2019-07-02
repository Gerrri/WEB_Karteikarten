package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var SessionNutzerID = ""

type tmp_b_home struct {
	Nutzername        string
	Nutzer            string
	NutzerEmail       string
	Lernkarten        string
	Karteien          string
	MeineKarteien     string
	NameVergeben      string
	EmailVergeben     string
	PasswortFalsch    string
	DatenschutzFalsch string
	ErstellteKarten   string
	ErstellteKarteien string
	MitgliedSeit      string
	Bild              string
	BildKlein         string
}

type tmp_L_lernen struct {
	//Menüleiste
	Nutzername    string
	Nutzer        string
	Lernkarten    string
	Karteien      string
	MeineKarteien string

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
	BildKlein    string
}

type tmp_nL_Karteikasten struct {
	Nutzername            string
	Karteien              string
	MeineKarteien         string
	Naturwissenschaften   []Karteikasten
	Sprachen              []Karteikasten
	Gesellschaft          []Karteikasten
	Wirtschaft            []Karteikasten
	Geisteswissenschaften []Karteikasten
	Sonstige              []Karteikasten
	BildKlein             string
}

type tmp_L_MeineKarteikaesten struct {
	Nutzername                string
	Karteien                  string
	MeineKarteien             string
	GespeicherteKarteikaesten []Karteikasten
	MeineKarteikaesten        []Karteikasten
	KastenID                  string
	KartenID                  int
	DelKastenID               string
	BildKlein                 string
}

type tmp_L_modkarteikasten1 struct {
	Nutzername            string
	Karteien              string
	MeineKarteien         string
	SwitchName            string
	AlleKarten            []Karte
	Wiederholungen        []int
	AktuelleKarte         Karte
	AktuellerKarteikasten Karteikasten
	//von aktueller Karte
	Titel   string
	Frage   string
	Antwort string

	//aktueller Kasten
	KastenID  string
	KartenID  int
	BildKlein string
}

/* ######################   not logged in Pages   ###################### */

func NL_Home(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()
		var nutzername = r.FormValue("nutzername")
		var passwort = r.FormValue("passwort")
		var nutzer = GetAlleNutzer()

		var isExecuted bool = false

		for _, arr := range nutzer {
			if arr.Name == nutzername && arr.Name != "" {
				if arr.Passwort == passwort {
					SessionNutzerID = arr.DocID
					//fmt.Println(SessionNutzerID)
					r.Method = ""
					isExecuted = true
					//http.Post("http://localhost/l_meinekarteikaesten", "", nil)
					http.Redirect(w, r, "./l_meinekarteikaesten", http.StatusSeeOther)
					break

				}
			}
		}

		if !isExecuted {
			p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
			t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in_popup.html")

			t.ExecuteTemplate(w, "layout", p)
		}
	} else {
		//fmt.Println("Joooo Broo")
		p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
		t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in.html")

		t.ExecuteTemplate(w, "layout", p)
	}
}

func NL_Home_popup(w http.ResponseWriter, r *http.Request) {

	p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/b_home.html", "./templates/nL_not_logged_in_popup.html")

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
		var geladen bool = true
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
			SessionNutzerID = AddNutzer(hinzufuegen)
			geladen = false
			http.Redirect(w, r, "./l_meinekarteikaesten", http.StatusSeeOther)
		}

		if geladen {
			p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz()), NameVergeben: "false", EmailVergeben: "false", PasswortFalsch: "false", DatenschutzFalsch: "false"}
			if vorhanden_benutzer {
				p.NameVergeben = "true"
			}
			if vorhanden_mail {
				p.EmailVergeben = "true"
			}
			if passwortUngleich {
				p.PasswortFalsch = "true"
			}
			if datenschutz != "on" {
				p.DatenschutzFalsch = "true"
			}
			t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_registrieren.html")

			t.ExecuteTemplate(w, "layout", p)
		}

	} else {
		p := tmp_b_home{Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz()), NameVergeben: "false", EmailVergeben: "false", PasswortFalsch: "false", DatenschutzFalsch: "false"}
		t, _ := template.ParseFiles("./templates/nL_not_logged_in.html", "./templates/nL_registrieren.html")

		t.ExecuteTemplate(w, "layout", p)
	}

}

/* ######################   logged in Pages   ###################### */
func L_Home(w http.ResponseWriter, r *http.Request) {
	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}

	p := tmp_b_home{BildKlein: link, Nutzername: GetNutzerById(SessionNutzerID).Name, Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/b_home.html", "./templates/L_logged_in.html")
	p.BildKlein = link
	t.ExecuteTemplate(w, "layout", p)
}

func L_karteikaesten(w http.ResponseWriter, r *http.Request) {
	data := tmp_nL_Karteikasten{
		Nutzername:            GetNutzerById(SessionNutzerID).Name,
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien:         strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),
		Naturwissenschaften:   []Karteikasten{},
		Sprachen:              []Karteikasten{},
		Gesellschaft:          []Karteikasten{},
		Wirtschaft:            []Karteikasten{},
		Geisteswissenschaften: []Karteikasten{},
		Sonstige:              []Karteikasten{},
	}

	kategorie := ""

	//POST Dropdown
	//Post auswertung
	if r.Method == "POST" {

		r.ParseForm()
		kategorie = r.FormValue("kategorie")
		fmt.Println("kategorie: ", kategorie)
		//Karteikästen nach Kategorien Laden

	}

	kk := []Karteikasten{}

	if kategorie == "" {
		kk = GetAlleKarteikaestenOeffentlich()
	} else {
		kk = SelectKarteikaestenByKategorie(GetAlleKarteikaestenOeffentlich(), kategorie)
	}

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

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_karteikaesten.html")
	data.BildKlein = link
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
		Nutzername:    GetNutzerById(SessionNutzerID).Name,
		Nutzer:        strconv.Itoa(GetNutzeranz()),
		Lernkarten:    strconv.Itoa(GetKartenAnz()),
		Karteien:      strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),

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

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_aufdecken.html")
	data.BildKlein = link
	t.ExecuteTemplate(w, "layout", data)
}

func L_lernen(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	//Kastenid und Kartenid auslesen
	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])

	var kasten = GetKarteikastenByid(Kastenid)
	var karte = kasten.Karten[Kartenid]

	//Kasten zu den Gespeicherten Kästen Hinzufügen, Wenn noch nicht vorhanden.
	AddKK2NutzerGespeichert(kasten, GetNutzerById(SessionNutzerID))

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
		Nutzername:    GetNutzerById(SessionNutzerID).Name,
		Nutzer:        strconv.Itoa(GetNutzeranz()),
		Lernkarten:    strconv.Itoa(GetKartenAnz()),
		Karteien:      strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),

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
	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}

	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_lernen.html")
	data.BildKlein = link
	t.ExecuteTemplate(w, "layout", data)
}

func L_meinekarteikaesten_popup(w http.ResponseWriter, r *http.Request) {
	data := tmp_L_MeineKarteikaesten{
		Nutzername:                GetNutzerById(SessionNutzerID).Name,
		Karteien:                  strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien:             strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),
		DelKastenID:               "",
		GespeicherteKarteikaesten: []Karteikasten{},
		MeineKarteikaesten:        []Karteikasten{},
	}

	var query = r.URL.Query()

	//Kastenid auslesen
	data.DelKastenID = (query["Kasten"])[0]

	nutzer := GetNutzerById(SessionNutzerID) //muss noch dynamisch gehlot werden

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

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten_popup.html")
	data.BildKlein = link
	t.ExecuteTemplate(w, "layout", data)

}

func L_meinekarteikaesten(w http.ResponseWriter, r *http.Request) {

	data := tmp_L_MeineKarteikaesten{
		Nutzername:                GetNutzerById(SessionNutzerID).Name,
		Karteien:                  strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien:             strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),
		GespeicherteKarteikaesten: []Karteikasten{},
		MeineKarteikaesten:        []Karteikasten{},
	}

	nutzer := GetNutzerById(SessionNutzerID) //muss noch dynamisch gehlot werden

	titel := ""
	beschreibung := ""
	kategorie := ""
	kategorieFilter := ""
	radio := ""
	if r.Method == "POST" {

		//POST Dropdown
		//Post auswertung
		if r.FormValue("kategorieFilter") != "" {
			r.ParseForm()
			kategorieFilter = r.FormValue("kategorieFilter")
			fmt.Println("kategorieFilter: ", kategorieFilter)

			//Karteikästen nach Kategorien Laden

		} else if r.FormValue("answer") == "" {
			fmt.Println("Löschen")

			var query = r.URL.Query()

			//Kastenid auslesen
			kID := (query["KastenID"])[0]

			//Löschen kk
			DeleteKarteikastenByID(kID)

		} else {

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
		}
	}

	for _, element := range nutzer.ErstellteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(SessionNutzerID)))

		if temp_kk.Kategorie == kategorieFilter || temp_kk.Unterkategorie == kategorieFilter {
			data.MeineKarteikaesten = append(data.MeineKarteikaesten, temp_kk)
		} else if kategorieFilter == "" {
			data.MeineKarteikaesten = append(data.MeineKarteikaesten, temp_kk)
		}

	}

	for _, element := range nutzer.GelernteKarteien {
		temp_kk := GetKarteikastenByid(element)
		temp_kk.FortschrittP = int(GetKarteikastenFortschritt(temp_kk, GetNutzerById(SessionNutzerID)))

		if temp_kk.Kategorie == kategorieFilter || temp_kk.Unterkategorie == kategorieFilter {
			data.GespeicherteKarteikaesten = append(data.GespeicherteKarteikaesten, temp_kk)
		} else if kategorieFilter == "" {
			data.GespeicherteKarteikaesten = append(data.GespeicherteKarteikaesten, temp_kk)
		}
	}

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinekarteikaesten.html")
	data.BildKlein = link
	t.ExecuteTemplate(w, "layout", data)
}

func L_meinProfil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Ich bin hier")
		r.ParseForm()
		var email = r.FormValue("email")
		var passwort = r.FormValue("passwort")
		var passwortNeu = r.FormValue("passwort_neu")
		var passwortNeuWdhl = r.FormValue("passwort_neuWdhl")

		var nutzer = GetNutzerById(SessionNutzerID)
		var alleNutzer = GetAlleNutzer()

		var vorhandenMail bool = false
		var passwortNichtGleich bool = false
		for _, arr := range alleNutzer {
			if arr.EMail == email {
				vorhandenMail = true
			}
		}

		if passwortNeu != passwortNeuWdhl {
			passwortNichtGleich = true
		}

		if vorhandenMail {
			//Fehlermeldung
		} else {
			if passwortNichtGleich {
				//Fehlermeldgung
			} else {
				if passwort == nutzer.Passwort {

					if passwortNeu != "" {
						nutzer.Passwort = passwortNeu
					}
					if email != "" {
						nutzer.EMail = email
					}
					fmt.Println(nutzer)
					UpdateNutzer(nutzer)
				}
			}
		}

	}

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	p := tmp_b_home{Nutzername: GetNutzerById(SessionNutzerID).Name, Bild: GetNutzerById(SessionNutzerID).Bild, MitgliedSeit: GetNutzerById(SessionNutzerID).MitgliedSeit, ErstellteKarteien: strconv.Itoa(len(GetNutzerById(SessionNutzerID).ErstellteKarteien)), NutzerEmail: GetNutzerById(SessionNutzerID).EMail, Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil.html")
	p.BildKlein = link
	t.ExecuteTemplate(w, "layout", p)
}

func L_meinProfil_popup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		DeleteNutzer(SessionNutzerID)
		r.Method = ""
		http.Redirect(w, r, "./nl_home", http.StatusSeeOther)
	}

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	p := tmp_b_home{Nutzername: GetNutzerById(SessionNutzerID).Name, Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil_popup.html")
	p.BildKlein = link
	t.ExecuteTemplate(w, "layout", p)
}

func L_meinProfil_popup_pic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var link = r.FormValue("link")
		var nutzer = GetNutzerById(SessionNutzerID)
		if link == "Hund" || link == "1" {
			nutzer.Bild = "https://svgsilh.com/svg/294256.svg"
		} else if link == "" {
			nutzer.Bild = "/icons/Mein-Profil_black.svg"
		} else if link == "Katze" || link == "2" {
			nutzer.Bild = "https://svgsilh.com/svg/2570357.svg"
		} else if link == "Ente" {
			nutzer.Bild = "/icons/Ente.svg"
		} else if link == "Eule" {
			nutzer.Bild = "/icons/Eule.svg"
		} else if link == "Grun" {
			nutzer.Bild = "/icons/Grun.svg"
		} else if link == "Sessel" {
			nutzer.Bild = "/icons/Sessel.svg"
		} else if link == "Zone30" {
			nutzer.Bild = "https://upload.wikimedia.org/wikipedia/commons/e/eb/Zeichen_274.1_-_Beginn_einer_Tempo_30-Zone%2C_StVO_2013.svg"
		} else {
			nutzer.Bild = link
		}

		UpdateNutzer(nutzer)

		http.Redirect(w, r, "http://localhost/l_meinProfil", http.StatusSeeOther)
	}

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	p := tmp_b_home{Nutzername: GetNutzerById(SessionNutzerID).Name, Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_meinProfil_popup_pic.html")
	p.BildKlein = link
	t.ExecuteTemplate(w, "layout", p)
}

func L_modkarteikasten1(w http.ResponseWriter, r *http.Request) {
	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	p := tmp_b_home{Nutzername: GetNutzerById(SessionNutzerID).Name, Nutzer: strconv.Itoa(GetNutzeranz()), Lernkarten: strconv.Itoa(GetKartenAnz()), MeineKarteien: strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)), Karteien: strconv.Itoa(GetKarteikastenAnz())}
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_modkarteikasten1.html")

	p.BildKlein = link
	t.ExecuteTemplate(w, "layout", p)
}

// NutzerID?
func L_modkarteikasten2(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	var Kastenid = (query["Kasten"])[0]
	var Kartenid, _ = strconv.Atoi((query["Karte"])[0])

	if query["Kasten"] != nil && query["Switch"] != nil {
		ToggleKarteikastenSichtbarkeit(Kastenid)
	}

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

	var karte = GetKarteikastenByid(Kastenid).Karten[Kartenid]

	data := tmp_L_modkarteikasten1{
		Nutzername:            GetNutzerById(SessionNutzerID).Name,
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien:         strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),
		AktuellerKarteikasten: Karteikasten{},
		AlleKarten:            []Karte{},
		SwitchName:            "",

		Titel:   karte.Titel,
		Frage:   karte.Frage,
		Antwort: karte.Antwort,

		KastenID: Kastenid,
		KartenID: Kartenid,
	}

	data.AktuellerKarteikasten = GetKarteikastenByid(Kastenid)
	data.AktuellerKarteikasten.FortschrittP = int(GetKarteikastenFortschritt(GetKarteikastenByid(Kastenid), GetNutzerById(SessionNutzerID)))

	if data.AktuellerKarteikasten.Sichtbarkeit == "Öffentlich" {
		data.SwitchName = "Privaten stellen!"
	} else {
		data.SwitchName = "Veröffentlichen!"
	}

	for i, element := range data.AktuellerKarteikasten.Karten {
		data.AlleKarten = append(data.AlleKarten, element)
		data.AlleKarten[i].Num = i + 1
		data.AlleKarten[i].Index = i
	}

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	data.BildKlein = link
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
		Nutzername:            GetNutzerById(SessionNutzerID).Name,
		Karteien:              strconv.Itoa(GetKarteikastenAnz()),
		MeineKarteien:         strconv.Itoa(GetKarteikastenAnzGespeicherte(SessionNutzerID)),
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

	var link = ""
	if GetNutzerById(SessionNutzerID).Bild == "/icons/Mein-Profil_black.svg" {
		link = "/icons/Mein-Profil.svg"
	} else {
		link = GetNutzerById(SessionNutzerID).Bild
	}
	data.BildKlein = link
	t, _ := template.ParseFiles("./templates/L_logged_in.html", "./templates/L_showKarteikarten.html")
	t.ExecuteTemplate(w, "layout", data)
}
