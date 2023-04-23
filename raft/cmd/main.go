package main

import (
	"log"
	"net/http"
)

func main() {
	g := LoadEngine()
	server := http.Server{
		addr: ":9090",
		mux:  g,
	}
	log.Fatal()
}
