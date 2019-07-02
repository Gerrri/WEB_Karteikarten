package controller

import (
	"couchdb"
	"encoding/json"
	"fmt"
	"strconv" //strconv.Itoa -> int to string
)

// Structs
type Nutzer struct {
	DocID             string `json:"_id"`
	DocRev            string `json:"_rev"`
	TYP               string
	Name              string
	EMail             string
	Passwort          string
	ErstellteKarteien []string
	GelernteKarteien  []string
}

type Karte struct {
	Num        int
	Index      int
	Titel      string
	Frage      string
	Antwort    string
	NutzerFach string
}

type Fortschritt struct {
	ID           string
	Wiederholung []int
}

type Karteikasten struct {
	DocID          string `json:"_id"`
	DocRev         string `json:"_rev"`
	TYP            string
	NutzerID       string
	Sichtbarkeit   string
	Kategorie      string
	Unterkategorie string
	Titel          string
	Anzahl         int
	Beschreibung   string
	Karten         []Karte
	Fortschritt    []Fortschritt
	FortschrittP   int
}

// Gibt die DB zurück (wenn nicht vorhanden = nil)
func GetDB() (d *couchdb.Database) {
	a, b := couchdb.NewDatabase(couchdb.DefaultBaseURL + "/web")
	var err error

	if b != nil {
		fmt.Print(b)
	} else {
		err = a.Available()
		if err == nil {
			//fmt.Println("DB is available")
			return a
		}
	}

	return nil
}

// ############################### START Kartei Methoden ############################### //
func GetKartenAnz() (anz int) {
	kk := GetAlleKarteikaesten()

	anz = 0

	for _, element := range kk {
		anz += len(element.Karten)
	}

	return anz
}

func GetKarteikastenFortschritt(k Karteikasten, nutzer Nutzer) (fortschritt float64) {
	fortschritt = 0
	var zaehler = 0
	xgesamt := len(k.Karten)

	if xgesamt == 0 {
		return 0
	}

	for n := 0; n < 4; n++ {
		zaehler += n * GetKarteikartenAnzByFach(k, n, nutzer)
	}

	fortschritt = float64(zaehler) / float64(4*float64(xgesamt)) * 100

	return fortschritt
}

func GetKarteikastenWiederholungArr(k Karteikasten, nutzer Nutzer) (i []int) {

	for _, element := range k.Fortschritt {
		if element.ID == nutzer.DocID {
			for _, wd := range element.Wiederholung {
				i = append(i, wd)
			}
			return i
		} else {
			return nil
		}
	}
	return nil

}

func GetKarteikartenAnzByFach(k Karteikasten, fach int, n Nutzer) (anz int) {
	var anzahl_fach = 0
	var wd = []int{}

	// Wiederholung im Fortschritt von Nutzer raussuchen
	wd = GetKKWiederholungenByNutzer(k, n)

	//Fachnummeranzahl entsprechend fach hochzählen
	for _, fachNr := range wd {
		if fachNr == fach {
			anzahl_fach++
		}
	}

	return anzahl_fach
}

// ############################### Ende Kartei Methoden ################################ //

func GetKKWiederholungenByNutzer(k Karteikasten, n Nutzer) (wd []int) {
	var Wiederholungen = []int{}

	//fmt.Println("NutzerID: ", n.DocID)
	for _, fort := range k.Fortschritt {

		//fmt.Println("FortID: ", fort.ID)
		if fort.ID == n.DocID {

			for _, aktwd := range fort.Wiederholung {
				Wiederholungen = append(Wiederholungen, aktwd)
			}

		}
	}

	return Wiederholungen
}

// ############################### START Karteikasten Methoden ############################### //

func UpdateKarteikastenKarte(KastenID string, KartenID int, n Nutzer, Richtig bool) {
	var db *couchdb.Database = GetDB()

	//func (db *Database) Save(doc interface{}, id string, rev string) (string, error)

	kk := GetKarteikastenByid(KastenID)
	wd := GetKKWiederholungenByNutzer(kk, n)
	//k := kk.Karten[KartenID]

	//Richtig
	if Richtig == true {
		//nur wenn Fortschritt kleriner 4 ++

		//damit es beim zurückspringen nicht zu "out of bounce" kommt
		if KartenID == -1 {
			KartenID = len(kk.Karten) - 1
		}

		if wd[KartenID] < 4 {
			wd[KartenID]++
		}
	}

	if Richtig == false {
		//nur wenn Fortschritt grüßer 0 --

		//damit es beim zurückspringen nicht zu "out of bounce" kommt
		if KartenID == -1 {
			KartenID = len(kk.Karten) - 1
		}

		if wd[KartenID] > 0 {
			wd[KartenID]--
		}
	}

	//fmt.Println("")

	//fmt.Println("id: ", kk.DocID)
	//fmt.Println("Rev: ", kk.DocRev)

	//altes Löschen & neues rein
	db.Set(kk.DocID, kk2Map(kk))

}

func UpdateKarteikarte(KastenID string, KartenID int, titel string, frage string, antwort string) {

	var db *couchdb.Database = GetDB()
	kk := GetKarteikastenByid(KastenID)

	kk.Karten[KartenID].Titel = titel
	kk.Karten[KartenID].Frage = frage
	kk.Karten[KartenID].Antwort = antwort

	db.Set(kk.DocID, kk2Map(kk))

}

func GetKarteikastenAnz() (anz int) {
	return len(GetAlleKarteikaesten())
}

func GetKarteikastenAnzGespeicherte(NutzerID string) (anz int) {

	return len(GetNutzerById(NutzerID).GelernteKarteien) + len(GetNutzerById(NutzerID).ErstellteKarteien)

}

func GetAlleKarteikaestenPrivat(nutzer Nutzer) (kk []Karteikasten) {
	allekk := GetAlleKarteikaesten()

	for _, element := range allekk {
		if element.Sichtbarkeit == "Privat" && element.NutzerID == nutzer.DocID {
			kk = append(kk, element)
		}
	}

	return kk
}

func GetAlleKarteikaestenOeffentlich() (kk []Karteikasten) {
	allekk := GetAlleKarteikaesten()

	for _, element := range allekk {
		if element.Sichtbarkeit == "Öffentlich" {
			kk = append(kk, element)
		}
	}

	return kk
}

func GetAlleKarteikaesten() (kk []Karteikasten) {
	var db *couchdb.Database = GetDB()

	var inmap []map[string]interface{}

	inmap, err := db.QueryJSON(`
	{
	  "selector": {
		"TYP": "Karteikasten"
	  }
	}`)

	for _, element := range inmap {
		var in = mapToJSON(element)

		var temp_kk = Karteikasten{}
		if err == nil {
			json.Unmarshal([]byte(in), &temp_kk)

			kk = append(kk, temp_kk)
		} else {
			fmt.Println(err)
		}
	}

	//for _, element := range kk {
	//	TerminalOutKarteikasten(element)
	//}

	return kk
}

func GetKarteikastenByid(id string) (k Karteikasten) {

	kk := GetAlleKarteikaesten()

	for _, element := range kk {
		if element.DocID == id {
			return element
		}
	}

	return k
}

func DeleteKarteikastenByID(id string) {
	var db *couchdb.Database = GetDB()
	change := false

	//Kasten auspflegen (Nutzer)

	allenutzer := GetAlleNutzer()

	for _, nutzer := range allenutzer {

		a := nutzer.ErstellteKarteien
		//Erstellete Karteien Löschen
		for i, ide := range nutzer.ErstellteKarteien {

			if ide == id {
				//ID Löschen
				// Remove the element at index i from a.
				nutzer.ErstellteKarteien = append(a[:i], a[i+1:]...)
				change = true

				db.Delete(id)
			}
		}

		a = nutzer.GelernteKarteien
		//Gespeicherte Karteien Löschen
		for i, ide := range nutzer.GelernteKarteien {
			if ide == id {
				//ID Löschen
				// Remove the element at index i from a.
				nutzer.GelernteKarteien = append(a[:i], a[i+1:]...)
				change = true
			}
		}

		//Nutzer Updaten
		if change == true {
			db.Set(nutzer.DocID, nutzer2Map(nutzer))
		}

	}

}

func ToggleKarteikastenSichtbarkeit(KastenID string) (Sichrbarkeit string) {
	var db *couchdb.Database = GetDB()

	kk := GetKarteikastenByid(KastenID)

	//fmt.Println("Sichtbarket: ", kk.Sichtbarkeit)

	if kk.Sichtbarkeit == "Öffentlich" {
		kk.Sichtbarkeit = "Privat"
	} else {
		kk.Sichtbarkeit = "Öffentlich"
	}

	//fmt.Println("Sichtbarket: ", kk.Sichtbarkeit)

	db.Set(kk.DocID, kk2Map(kk))

	return kk.Sichtbarkeit
}

func AddKarteikasten(kk Karteikasten, nutzer Nutzer) error {
	var db *couchdb.Database = GetDB()

	f := Fortschritt{}
	f.ID = nutzer.DocID
	kk.NutzerID = nutzer.DocID

	kk.Fortschritt = append(kk.Fortschritt, f)
	// Convert Todo suct to map[string]interface as required by Save() method
	KarteiK := kk2Map(kk)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(KarteiK, "_id")
	delete(KarteiK, "_rev")
	delete(KarteiK, "FortschrittP")

	// Add todo to DB
	id, _, err := db.Save(KarteiK, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	//Karte hinzufügen
	AddKarteikarte(id, "Meine Erste Karteikarte", "Schreibe hier deine Frage", "...und hier die Antwort :)")

	//Update Nutzer

	AddKKtoNutzer(nutzer, GetKarteikastenByid(id))
	//db.Save(KarteiK, nil)

	return err
}

func AddKarteikarte(KastenID string, titel string, frage string, antwort string) {
	var db *couchdb.Database = GetDB()

	k := Karte{}
	k.Titel = titel
	k.Frage = frage
	k.Antwort = antwort

	kk := GetKarteikastenByid(KastenID)
	for i := 0; i < len(kk.Fortschritt); i++ {
		kk.Fortschritt[i].Wiederholung = append(kk.Fortschritt[i].Wiederholung, 0)
	}

	kk.Karten = append(kk.Karten, k)

	fmt.Println("Neue Karte hinzufügen ...")

	kkmap := kk2Map(kk)

	delete(kkmap, "NutzerFach")
	delete(kkmap, "Num")
	delete(kkmap, "Index")

	db.Set(kk.DocID, kkmap)
}

func DelKarteikarteByID(KastenID string, KartenID int) {
	var db *couchdb.Database = GetDB()

	kk := GetKarteikastenByid(KastenID)

	//Karten werden neu befüllt
	newKarten := []Karte{}
	for i, Karte := range kk.Karten {
		if i != KartenID {
			newKarten = append(newKarten, Karte)
		}
	}

	//Wiederholungen werden neu befüllt
	newWiederholung := []int{}

	for i, _ := range kk.Fortschritt {
		for j, Wiederholung := range kk.Fortschritt[i].Wiederholung {
			if j != KartenID {
				newWiederholung = append(newWiederholung, Wiederholung)
			}
		}

		kk.Fortschritt[i].Wiederholung = newWiederholung
		newWiederholung = []int{}
	}

	kk.Karten = newKarten

	//fmt.Println("kk: ", kk)

	db.Set(kk.DocID, kk2Map(kk))
}

func TerminalOutKarteikasten(k Karteikasten) {
	fmt.Println("############# KARTEIKASTEN ##############")
	fmt.Println("id : " + k.DocID)
	fmt.Println("NutzerID : " + k.NutzerID)
	fmt.Println("Oeffentlich : " + k.Sichtbarkeit)
	fmt.Println("Kategorie : " + k.Kategorie)
	fmt.Println("Unterkategorie : " + k.Unterkategorie)
	fmt.Println("Titel : " + k.Titel)
	fmt.Println("Anzahl : " + strconv.Itoa(k.Anzahl))
	fmt.Println("Beschreibung : " + k.Beschreibung)
	fmt.Println("#########################################")
}

// ############################### ENDE Karteikasten Methoden ############################### //

// ############################### START Nutzer Methoden ############################### //

//Wenn nicht bereits vorhanden
func AddKK2NutzerGespeichert(kk Karteikasten, n Nutzer) {
	var db *couchdb.Database = GetDB()
	var vorhanden = false

	// #######  Start Update KK Fortschitt
	//Add Fortschritt to Nutzer

	for _, fort := range kk.Fortschritt {
		if fort.ID == n.DocID {
			vorhanden = true
		}
	}

	if vorhanden == false {
		f := Fortschritt{}
		f.ID = n.DocID

		for i := 0; i < len(kk.Karten); i++ {
			f.Wiederholung = append(f.Wiederholung, 0)
		}

		//Fortschritt f KK hinzufügen
		kk.Fortschritt = append(kk.Fortschritt, f)
	}

	db.Set(kk.DocID, kk2Map(kk))
	// #######  Ende Update KK Fortschitt

	vorhanden = false

	for _, id := range n.GelernteKarteien {
		if id == kk.DocID {
			vorhanden = true
		}
	}

	if vorhanden == false && n.DocID != kk.NutzerID {
		n.GelernteKarteien = append(n.GelernteKarteien, kk.DocID)
	}

	db.Set(n.DocID, nutzer2Map(n))
}

//Wenn nicht vorhanden ID = -1
func GetNutzerById(id string) (n Nutzer) {

	var arr = GetAlleNutzer()

	for _, nutzer := range arr {
		if nutzer.DocID == id {
			return nutzer
		}
	}

	n = Nutzer{}
	n.DocID = "null"
	fmt.Println("N: ", n)
	return n
}

//-1 = db not da
//-2 = abfrage nicht möglich
func GetNutzeranz() (anz int) {
	var n = GetAlleNutzer()

	return len(n)

	return -1
}

func AddKKtoNutzer(n Nutzer, kk Karteikasten) {
	var db *couchdb.Database = GetDB()

	n.ErstellteKarteien = append(n.ErstellteKarteien, kk.DocID)

	db.Set(n.DocID, nutzer2Map(n))
}

func AddNutzer(n Nutzer) (id string) {
	var db *couchdb.Database = GetDB()
	var nutzermap = nutzer2Map(n)
	delete(nutzermap, "_id")
	delete(nutzermap, "_rev")

	id, _, err := db.Save(nutzermap, nil)
	fmt.Println(id, err)

	return id
}

func GetAlleNutzer() (n []Nutzer) {
	var db *couchdb.Database = GetDB()

	inmap, err := db.QueryJSON(`
	{
		"selector": {
		"TYP": "nutzer"
		}
	}`)

	for _, element := range inmap {

		var in = mapToJSON(element)

		var temp_an = Nutzer{}
		if err == nil {
			json.Unmarshal([]byte(in), &temp_an)

			n = append(n, temp_an)

		} else {
			fmt.Println(err)
		}
	}

	return n
}

func TerminalOutNutzer(n Nutzer) {
	fmt.Println("ID 		: " + n.DocID)
	fmt.Println("Name 		: " + n.Name)
	fmt.Println("Email 		: " + n.EMail)
	fmt.Println("Passwort 	: " + n.Passwort)
}

// ############################### ENDE Nutzer Methoden ############################### //

func mapToJSON(inMap map[string]interface{}) (s string) {
	var b []byte

	b, err := json.Marshal(inMap)
	jsonString := string(b)

	//Error Output
	if err != nil {
		fmt.Print("JSON Convertion Error: ")
		fmt.Println(err)
	}

	return jsonString
}

func kk2Map(kk Karteikasten) map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(kk)
	json.Unmarshal(tJSON, &doc)

	return doc
}

func nutzer2Map(n Nutzer) map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(n)
	json.Unmarshal(tJSON, &doc)

	return doc
}

func k2Map(k Karteikasten) map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(k)
	json.Unmarshal(tJSON, &doc)

	return doc
}
