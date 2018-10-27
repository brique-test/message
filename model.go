package main

// 메시지를 주고받을 때 사용할 구조체 정의

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
