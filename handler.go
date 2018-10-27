package main

import (
	"html/template"
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}
