package main

import(
	"net/http"
	"/home/darthbaluu/Dokumente/Uni/6. Semester/WEB/WEB_Karteikarten/BrainTrainGo/controller/controller"
)

func main(){
	//eingeloggt
	http.HandleFunc("/eingeloggt_startseite", controller.eingeloggt_start)
	http.HandleFunc("/eingeloggt_karteikasten", controller.eingeloggt_karteikasten)
	http.HandleFunc("/eingeloggt_meinekarteien", controller.eingeloggt_meineKarteien)
	http.HandleFunc("/eingeloggt_meinprofil", controller.eingeloggt_meinProfil)
	http.HandleFunc("/eingeloggt_karteierstellen_01", controller.eingeloggt_karteiErstellen_01)
	http.HandleFunc("/eingeloggt_karteierstellen_02", controller.eingeloggt_karteiErstellen_02)
	http.HandleFunc("/eingeloggt_karteikastenansehen", controller.eingeloggt_karteikastenAnsehen)
	http.HandleFunc("/eingeloggt_lernen_01", controller.eingeloggt_lernen_01)
	http.HandleFunc("/eingeloggt_lernen_02", controller.eingeloggt_lernen_02)
	http.HandleFunc("/eingeloggt_profilloeschen", controller.eingeloggt_profilLoeschen)

	//ausgeloggt
	http.HandleFunc("/ausgeloggt_startseite", controller.ausgeloggt_start)
	http.HandleFunc("/ausgeloggt_karteikasten", controller.ausgeloggt_karteikasten)
	http.HandleFunc("/ausgeloggt_registrieren", controller.ausgeloggt_registrieren)

	server := http.Server{
		Addr: ":80",
	}

	server.ListenAndServe()

}

