package main

import (
	"flag"
	"net/http"

	"github.com/milosz94/web-go/api"
)

func main() {

	listenAddr := flag.String("listenaddr", ":8080", "Address for http listening")
	flag.Parse()

	http.HandleFunc("/user", api.HandleGetUser)
	http.HandleFunc("/account", api.HandleGetAccount)

	http.ListenAndServe(*listenAddr, nil)
}
