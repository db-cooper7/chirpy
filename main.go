package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const rootFilePath = "."

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(rootFilePath)))

	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
