package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiartweb/internal/asciiart"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")       // value to be changed to what input side name identifier has
	// address := r.FormValue("address") // value of output?
	fmt.Fprintf(w, "name = %s\n", name)
	// fmt.Fprintf(w, "address = %s\n", address)
	fmt.Println(name)
	// fmt.Printf( "address = %s\n", address)
	asciiart.AsciiArt()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found dumbass", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("../../static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
