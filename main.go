package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error. Error: %v", err)
	}
	fmt.Fprintf(w, "POST request successful!\n")

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %v\t Email: %v", name, email)
}
