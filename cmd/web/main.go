package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":3000", "Server Port")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on %s", *port)

	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}
