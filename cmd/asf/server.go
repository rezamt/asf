package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s\n", r.Method, r.URL.Path)
		fmt.Fprintf(w, "{ \"path\" : %q}", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}