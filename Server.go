package main

import(
	"net/http"
	"./app/controller"
)

func main(){
	
	http.HandleFunc("/eingeloggt_startseite", controller.Eingeloggt_start)
	http.HandleFunc("/eingeloggt_karteikasten", controller.Eingeloggt_karteikasten)
	http.HandleFunc("/eingeloggt_meinekarteien", controller.Eingeloggt_meineKarteien)
	http.HandleFunc("/eingeloggt_meinprofil", controller.Eingeloggt_meinProfil)
	http.HandleFunc("/eingeloggt_karteierstellen_01", controller.Eingeloggt_karteiErstellen_01)
	http.HandleFunc("/eingeloggt_karteierstellen_02", controller.Eingeloggt_karteiErstellen_02)
	http.HandleFunc("/eingeloggt_karteikastenansehen", controller.Eingeloggt_karteikastenAnsehen)
	http.HandleFunc("/eingeloggt_lernen_01", controller.Eingeloggt_lernen_01)
	http.HandleFunc("/eingeloggt_lernen_02", controller.Eingeloggt_lernen_02)
	http.HandleFunc("/eingeloggt_profilloeschen", controller.Eingeloggt_profilLoeschen)

	
	http.HandleFunc("/ausgeloggt_startseite", controller.Ausgeloggt_start)
	http.HandleFunc("/ausgeloggt_karteikasten", controller.Ausgeloggt_karteikasten)
	http.HandleFunc("/ausgeloggt_registrieren", controller.Ausgeloggt_registrieren)

	//bereitstellung der statischen inhalte
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/favicons/", http.StripPrefix("/favicons/", http.FileServer(http.Dir("./static/favicons"))))
	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("./static/font"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("./static/icons"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	http.Handle("/logo/", http.StripPrefix("/logo/", http.FileServer(http.Dir("./static/logo"))))

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()

}

