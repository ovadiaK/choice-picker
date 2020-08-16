package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	// init templates
	templates = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// choice receiver
	// choice sender
	fmt.Println(templates)
	if err := templates.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		panic(err)
	}
}
