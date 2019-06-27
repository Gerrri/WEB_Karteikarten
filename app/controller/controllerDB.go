package controller

import (
	"couchdb"
	"encoding/json"
	"errors"
	"fmt"
	"strconv" //strconv.Itoa -> int to string
)

// Structs
type Nutzer struct {
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
	Nutzer []Nutzer
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
	ID           int
	Wiederholung []int
}

type Karteikasten struct {
	ID             int
	_id            string
	_rev           string
	NutzerID       int
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

	for n := 0; n < 4; n++ {
		zaehler += n * GetKarteikartenAnzByFach(k, n, nutzer)
	}

	fortschritt = float64(zaehler) / float64(4*float64(xgesamt)) * 100

	return fortschritt
}

func GetKarteikastenWiederholungArr(k Karteikasten, nutzer Nutzer) (i []int) {

	for _, element := range k.Fortschritt {
		if element.ID == nutzer.ID {
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
	for index, _ := range k.Karten {
		if k.Fortschritt[n.ID].Wiederholung[index] == fach {
			anzahl_fach++
		}
	}

	return anzahl_fach
}

// ############################### Ende Kartei Methoden ################################ //

// ############################### START Karteikasten Methoden ############################### //

func UpdateKarteikastenKarte(KastenID int, KartenID int, NutzerID int, Richtig bool) {
	var db *couchdb.Database = GetDB()

	//func (db *Database) Save(doc interface{}, id string, rev string) (string, error)

	kk := GetKarteikastenByid(KastenID)
	//k := kk.Karten[KartenID]
	NutzerID = 1
	NutzerID--

	//Richtig
	if Richtig == true {
		//nur wenn Fortschritt kleriner 4 ++

		//damit es beim zurückspringen nicht zu "out of bounce" kommt
		if KartenID == -1 {
			KartenID = len(kk.Karten) - 1
		}

		if kk.Fortschritt[NutzerID].Wiederholung[KartenID] < 4 {
			kk.Fortschritt[NutzerID].Wiederholung[KartenID]++
		}
	}

	if Richtig == false {
		//nur wenn Fortschritt grüßer 0 --
		if kk.Fortschritt[NutzerID].Wiederholung[KartenID] > 0 {
			kk.Fortschritt[NutzerID].Wiederholung[KartenID]--
		}
	}

	//altes Löschen & neues rein
	kk._id = "KK_2_das_kleine_1x1"
	kk._rev = "19-0e9dadc6109851185f21730706bbfe0b"
	db.DeleteDoc(kk2Map(kk))
	db.Save(kk2Map(kk), nil)
}

func GetKarteikastenAnz() (anz int) {
	return len(GetAlleKarteikaesten())
}

func GetAlleKarteikaestenPrivat(NutzerID int) (kk []Karteikasten) {
	allekk := GetAlleKarteikaesten()

	for _, element := range allekk {
		if element.Sichtbarkeit == "Privat" && element.NutzerID == NutzerID {
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

//Wenn nicht vorhanden ID = -1
func GetNutzerById(id int) (n Nutzer) {

	var arr, err = getNutzerArr()

	if err == nil {
		for _, n := range arr {
			if n.ID == id {
				return n
			}
		}
	}

	n = Nutzer{}
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

func getNutzerArr() (n []Nutzer, err error) {
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

func TerminalOutNutzer(n Nutzer) {
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

func kk2Map(kk Karteikasten) map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(kk)
	json.Unmarshal(tJSON, &doc)

	return doc
}
