package main

import (
	"BrainTrain/app/controller"
	"net/http"
)

func main() {

	//Pages Logged in
	http.HandleFunc("/l_home", controller.L_Home)
	http.HandleFunc("/l_changeKK", controller.L_changeKK)
	http.HandleFunc("/l_aufdecken", controller.L_aufdecken)
	http.HandleFunc("/l_karteikaesten", controller.L_karteikaesten)
	http.HandleFunc("/l_lernen", controller.L_lernen)
	http.HandleFunc("/l_meinekarteikaesten", controller.L_meinekarteikaesten)
	http.HandleFunc("/l_meinekarteikaesten_popup", controller.L_meinekarteikaesten_popup)
	http.HandleFunc("/l_meinProfil", controller.L_meinProfil)


	http.HandleFunc("/l_meinProfil_popup", controller.L_meinProfil_popup)
	http.HandleFunc("/l_meinProfil_popup_pic", controller.L_meinProfil_popup_pic)
	http.HandleFunc("/l_modkarteikasten1", controller.L_modkarteikasten1)
	http.HandleFunc("/l_modkarteikasten2", controller.L_modkarteikasten2)
	http.HandleFunc("/l_showKarteikarten", controller.L_showKarteikarten)

	//Pages not Logged in
	http.HandleFunc("/", controller.NL_Home)
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
		Addr: ":80",
	}

	server.ListenAndServe()

}
