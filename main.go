package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "./Public"
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/",http.FileServer(http.Dir(filepathRoot))))
	


	srv := &http.Server{
		Addr:  ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port %s/n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
