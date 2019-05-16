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

// ###############################  ############################### //

// Nutzer Methoden //

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

func nutzerTerminalOut(n nutzer) {
	fmt.Println("ID 		: " + strconv.Itoa(n.ID))
	fmt.Println("Vorname 	: " + n.Vorname)
	fmt.Println("Name 		: " + n.Name)
	fmt.Println("Email 		: " + n.EMail)
	fmt.Println("Passwort 	: " + n.Passwort)
}

// Nutzer Methoden //

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
