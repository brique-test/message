package main

import "github.com/gorilla/websocket"

var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
