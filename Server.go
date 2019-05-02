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

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()

}
