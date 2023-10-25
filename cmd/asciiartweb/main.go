package main

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"text/template"

	"asciiartweb/internal/asciiart"
)

var temp1 *template.Template

func init() {
	temp1 = template.Must(template.ParseFiles("../../static/404.html"))
}

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
	// err = parsedTemplate.Execute(w, fonts)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }

	// if name == "" || name >= string(32) || name <= string(128) {
	// 	http.ServeFile(w, r, "../../static/400.html")
	// 	return
	// }
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		helloHandler(w, r)
	}
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("../../static/400.html")
	parsedTemplate.Execute(w, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	temp1.ExecuteTemplate(w, "404.html", r)
	w.WriteHeader(404)
	parsedTemplate, _ := template.ParseFiles("../../static/404.html")
	parsedTemplate.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("../../static/500.html")
	parsedTemplate.Execute(w, nil)
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

	// if r.URL.Path != "/" {
	// 	return
	// }

}

func main() {
	http.HandleFunc("/", printHandler)
	http.HandleFunc("/ascii-art", formHandler)
	http.HandleFunc("/400", hiHandler) // 400 bad request
	http.HandleFunc("/404", helloHandler)
	http.HandleFunc("/500", errorHandler)
	http.HandleFunc("/w.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../static/w.css")
	})

	fmt.Printf("Starting server at http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
