package main

import (
	"log"
	"net/http"
)

func main()  {
	log.Println("Server start on port 8000")
	http.HandleFunc("/", HtmlHandler)
	http.ListenAndServe(":8000", nil)
}