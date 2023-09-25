package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.FPrintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.FPrintf(w, "POST request successful!")
	name := r.FormValue("name") // value to be changed to what input side name identifier has
	address := r.ForwardValue("address") // value 


}


func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello"{
http.Error(w, "404 not found", http.StatusNotFound)
return
}
if r.Method != "GET"{
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
	}
	fmt.FPrintf(w, "hello!")
}

func main(){

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListedAndSevre(":8080", nil); err != nil {
		log.Fatal(err)
	}
}