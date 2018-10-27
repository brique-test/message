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
	log.Println("New connection created")

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("conn: ", err)
			delete(clients, conn)
			break
		}

		broadcast <- msg
		log.Printf("[%v] %v\n", msg.Username, msg.Content)
	}
}

func MessageHandler() {
	for {
		msg := <-broadcast

		for conn := range clients {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("message: ", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}
