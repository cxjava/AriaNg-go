package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

func startServer(addr string) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(packr.New("AriaNg", "./AriaNg"))))
	log.Println("starting server on " + addr)
	log.Println(http.ListenAndServe(addr, nil))
}
