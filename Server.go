package main

import (
	"BrainTrain/app/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/test", controller.Test_site)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/nl_home", controller.NL_Home)

	http.HandleFunc("/nL_karteikaesten", controller.NL_karteikaesten)

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
