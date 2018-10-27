package main

import (
	"html/template"
	"log"
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func ConnHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Panicf("error: %v\n", err)
			delete(clients, conn)
			break
		}

		broadcast <- msg
	}
}

func MessageHandler()  {
	for {
		msg := <- broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Panicf("error: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}