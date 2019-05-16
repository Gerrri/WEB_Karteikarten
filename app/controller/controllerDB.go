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
	ID                    int
	Vorname               string
	Name                  string
	EMail                 string
	Passwort              string
	ErstellteKarteikarten []int
	GelernteKarteikarten  []int
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
	ID             string
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
			fmt.Println("DB is available")
			return a
		}
	}

	return nil
}

// ############################### START Karteikasten Methoden ############################### //
func GetKarteikastenByid(id string) (k Karteikasten) {

	var db *couchdb.Database = GetDB()

	//result, err = db.Get("nutzer", nil)
	var result, err = db.Get(id, nil)

	if err == nil {
		in := mapToJSON(result)

		kk := Karteikasten{}
		json.Unmarshal([]byte(in), &kk)
		return kk
	}

	kk := Karteikasten{}
	kk.ID = "-1"
	return kk
}

func TerminalOutKarteikasten(k Karteikasten) {
	fmt.Println("id : " + k.ID)
	fmt.Println("NutzerID : " + strconv.Itoa(k.NutzerID))
	fmt.Println("Oeffentlich : " + strconv.FormatBool(k.Oeffentlich))
	fmt.Println("Kategorie : " + k.Kategorie)
	fmt.Println("Unterkategorie : " + k.Unterkategorie)
	fmt.Println("Titel : " + k.Titel)
	fmt.Println("Anzahl : " + strconv.Itoa(k.Anzahl))
	fmt.Println("Beschreibung : " + k.Beschreibung)

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
