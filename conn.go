package main

import "github.com/gorilla/websocket"

// 커넥션 객체들을 모아두는 map 객체
// Key: WebSocket Conn 포인터
// Value: boolean
// Key 위치에 Conn 을 놓으면 Map에서 Conn을 추가하고 제거하기가 편리함
// Value는 쓰이고 있지 않음 (리팩토링 필요)
var clients = make(map[*websocket.Conn]bool)

// goroutine이 데이터를 주고받는 통로
// 타입은 Message 구조체
var broadcast = make(chan Message)

// WebSocket이 프로토콜을 Upgrade하도록 해주는 구조체
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
