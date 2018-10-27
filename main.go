package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server start on port 8000")
	// 웹브라우저로 접속하면 index.html을 전송하는 핸들러
	http.HandleFunc("/", HtmlHandler)
	// WebSocket으로 프로토콜을 변경해주는 커넥션 핸들러
	http.HandleFunc("/ws", ConnHandler)

	// goroutine으로 비동기적으로 실행되는 메시지 핸들러
	go MessageHandler()

	// 8000 포트로 서버 오픈
	err := http.ListenAndServe(":8000", nil)

	// 에러 처리
	if err != nil {
		log.Panicln("server error: ", err)
	}

	// 에러로 인해 Panic이 작동할 경우 main 함수를 되살리는 recover 함수
	// defer 키워드: Panic으로 인해 main 함수가 종료되기 직전에 실행되도록 함
	defer func() {
		s := recover()
		log.Println(s)
	}()
}
