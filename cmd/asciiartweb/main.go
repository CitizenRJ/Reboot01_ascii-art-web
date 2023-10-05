package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"asciiartweb/internal/asciiart"
)

type Fonts struct {
	Art    string
	Hidden string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	banner := r.FormValue("banner")
	fmt.Println("banner= " + banner)
	art := asciiart.AsciiArt(banner, name)
	fonts := Fonts{Art: art, Hidden: "false"}
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

func printHandler(w http.ResponseWriter, r *http.Request) {

	parsedTemplate, err := template.ParseFiles("../../static/index.html")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {

	http.HandleFunc("/", printHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/w.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../static/w.css")
	})

	fmt.Printf("Starting server at http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
