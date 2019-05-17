package controller

import (
	"couchdb"
	"encoding/json"
	"errors"
	"fmt"
	"strconv" //strconv.Itoa -> int to string
)

// Structs
type nutzer struct {
	ID                int
	Vorname           string
	Name              string
	EMail             string
	Passwort          string
	ErstellteKarteien []int
	GelernteKarteien  []int
}

type alleNutzer struct {
	_id    string
	_rev   string
	Nutzer []nutzer
}

type Karte struct {
	Titel   string
	Frage   string
	Antwort string
}

type Fortschritt struct {
	ID           int
	Wiederholung []int
}

type Karteikasten struct {
	ID             int
	_rev           string
	NutzerID       int
	Oeffentlich    bool
	Kategorie      string
	Unterkategorie string
	Titel          string
	Anzahl         int
	Beschreibung   string
	Karten         []Karte
	Fortschritt    []Fortschritt
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

// ############################### Ende Kartei Methoden ################################ //

// ############################### START Karteikasten Methoden ############################### //
func GetKarteikastenAnz() (anz int) {
	return len(GetAlleKarteikaesten())
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

func GetKarteikastenByid(id int) (k Karteikasten) {

	kk := GetAlleKarteikaesten()

	for _, element := range kk {
		if element.ID == id {
			return element
		}
	}

	return k
}

func TerminalOutKarteikasten(k Karteikasten) {
	fmt.Println("############# KARTEIKASTEN ##############")
	fmt.Println("id : " + strconv.Itoa(k.ID))
	fmt.Println("NutzerID : " + strconv.Itoa(k.NutzerID))
	fmt.Println("Oeffentlich : " + strconv.FormatBool(k.Oeffentlich))
	fmt.Println("Kategorie : " + k.Kategorie)
	fmt.Println("Unterkategorie : " + k.Unterkategorie)
	fmt.Println("Titel : " + k.Titel)
	fmt.Println("Anzahl : " + strconv.Itoa(k.Anzahl))
	fmt.Println("Beschreibung : " + k.Beschreibung)
	fmt.Println("#########################################")
}

// ############################### ENDE Karteikasten Methoden ############################### //

// ############################### START Nutzer Methoden ############################### //

//Wenn nicht vorhanden ID = -1
func GetNutzerById(id int) (n nutzer) {

	var arr, err = getNutzerArr()

	if err == nil {
		for _, n := range arr {
			if n.ID == id {
				return n
			}
		}
	}

	n = nutzer{}
	n.ID = -1
	return n
}

//-1 = db not da
//-2 = abfrage nicht möglich
func GetNutzeranz() (anz int) {
	var n, err = getNutzerArr()

	if err == nil {
		return len(n)
	} else {
		fmt.Println(err)
		return -2
	}

	return -1
}

func getNutzerArr() (n []nutzer, err error) {
	var db *couchdb.Database = GetDB()

	if db == nil {
		return nil, errors.New("Datenbank Verbindung nicht möglich!")
	}

	//Nutzer Wählen
	var result map[string]interface{}

	//result, err = db.Get("nutzer", nil)
	result, err = db.Get("nutzer", nil)

	in := mapToJSON(result)

	an := alleNutzer{}
	json.Unmarshal([]byte(in), &an)

	return an.Nutzer, nil

}

func TerminalOutNutzer(n nutzer) {
	fmt.Println("ID 		: " + strconv.Itoa(n.ID))
	fmt.Println("Vorname 	: " + n.Vorname)
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
