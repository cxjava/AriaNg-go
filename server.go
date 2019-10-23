package main

import (
	"log"
	"net/http"

	"github.com/cxjava/AriaNg-go/assets"
)

func startServer(addr string) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(assets.HTTP)))
	log.Println("starting server on " + addr)
	log.Println(http.ListenAndServe(addr, nil))
}
