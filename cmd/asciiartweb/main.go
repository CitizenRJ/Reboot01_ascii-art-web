package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"asciiartweb/internal/asciiart"
)

type Fonts struct {
	Art string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	banner := r.FormValue("banner")
	fmt.Println(banner) // value to be changed to what input side name identifier has
	// address := r.FormValue("address") // value of output?
	// if r.URL.Path != "/hello" {
	// http.FileServer(http.Dir("../../static/404.html"))
	// http.Error(w, "404 not found dumbass", http.StatusNotFound)
	// fmt.Fprintf(w, "name = %s\n", name)
	// fmt.Fprintf(w, "address = %s\n", address)
	// fmt.Println(name)
	// fmt.Printf( "address = %s\n", address)
	art := asciiart.AsciiArt(banner, name)
	fonts := Fonts{Art: art}
	parsedTemplate, err := template.ParseFiles("../../static/index.html")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	err = parsedTemplate.Execute(w, fonts)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.FileServer(http.Dir("../../static/404.html"))
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
