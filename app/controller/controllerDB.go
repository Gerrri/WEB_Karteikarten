package controller

import (
	"couchdb"
	"encoding/json"
	"fmt"
)

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

//-1 = db not da
//-2 = abfrage nicht möglich
func GetNutzeranz() (anz int) {
	var db *couchdb.Database = GetDB()

	if db == nil {
		return -1
	}

	//Nutzer Wählen
	var result map[string]interface{}
	var err error
	//result, err = db.Get("nutzer", nil)
	result, err = db.Get("nutzer", nil)

	var i int = 0

	if err == nil {

		//enc := json.NewEncoder(os.Stdout)
		//enc.Encode(result)

		var b []byte

		b, err := json.Marshal(result)
		jsonString := string(b)

		fmt.Println(jsonString)
		fmt.Println(err)

		i++

	} else {
		fmt.Println(err)
		return -2
	}

	return 0
}
