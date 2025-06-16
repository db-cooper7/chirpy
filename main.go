package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const rootFilePath = "."
	const logoFilePath = "./assets/logo.png"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(rootFilePath)))
	mux.Handle("/assets", http.FileServer(http.Dir(logoFilePath)))

	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
