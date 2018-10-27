package main

import (
	"html/template"
	"log"
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func main()  {
	log.Println("Server start on port 8000")
	http.HandleFunc("/", HtmlHandler)
	http.ListenAndServe(":8000", nil)
}