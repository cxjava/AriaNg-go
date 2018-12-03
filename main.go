package main

import (
	"flag"
	"net/http"
)

var (
	address = flag.String("a", ":18080", `bind address.`)
)

func init() {
	flag.Parse()
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(HTTP)))
	http.ListenAndServe(*address, nil)
}
