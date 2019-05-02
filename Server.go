package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test_site)
	http.HandleFunc("/login", login)
	http.HandleFunc("/nl_home", nL_Home)

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()

}
