package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server start on port 8000")
	http.HandleFunc("/", HtmlHandler)
	http.HandleFunc("/ws", ConnHandler)

	go MessageHandler()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		s := recover()
		log.Println(s)
	}()
}
