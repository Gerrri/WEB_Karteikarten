package main

import (
	"BrainTrain/app/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/l_home", controller.L_Home)
	http.HandleFunc("/l_aufdecken", controller.L_aufdecken)
	http.HandleFunc("/l_karteikaesten", controller.L_karteikaesten)
	http.HandleFunc("/l_lernen", controller.L_lernen)
	http.HandleFunc("/l_meinekarteikaesten", controller.L_meinekarteikaesten)
	http.HandleFunc("/nl_home", controller.NL_Home)
	http.HandleFunc("/nl_karteikaesten", controller.NL_karteikaesten)
	http.HandleFunc("/nl_registrieren", controller.NL_registrieren)

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
