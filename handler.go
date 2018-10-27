package main

import (
	"html/template"
	"log"
	"net/http"
)

// index.html을 전송하는 핸들러
func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	// Must 함수: 에러 처리를 미리 도와주는 함수
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

// WebSocket 프로토콜을 열어주는 핸들러
func ConnHandler(w http.ResponseWriter, r *http.Request) {
	// conn.go 에서 정의한 upgrader 구조체의 Upgrade 메서드 실행
	// http 프로토콜을 WebSocket 프로토콜로 switching 해줌
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln(err)
	}
	// 함수가 종료되기 직전에 커넥션을 닫는다
	defer conn.Close()

	// clients map에 새로 생성한 커넥션 등록
	// 모든 커넥션이 이곳에 저장되므로 메시지 제어를 편리하게 할 수 있음
	clients[conn] = true
	log.Println("New connection created")

	// 무한 루프
	// 커넥션 객체에 JSON 메시지가 들어오는 걸 감지할 때에만 작동
	for {
		// Message 타입의 객체 정의
		var msg Message
		// 커넥션에서 JSON 메시지를 읽어들여 msg 레퍼런스에 덮어씀
		err := conn.ReadJSON(&msg)
		// 에러 처리: 커넥션에서 에러가 발생하면 clients map에서 커넥션을 지우고 루프 종료
		if err != nil {
			log.Println("conn: ", err)
			delete(clients, conn)
			break
		}

		// 메시지를 정상적으로 읽었다면 msg를 broadcast 채널로 송신
		broadcast <- msg
		log.Printf("[%v] %v\n", msg.Username, msg.Content)
	}
}

func MessageHandler() {
	// 무한 루프
	// broadcast에 메시지가 들어오는 걸 감지할 때에만 작동
	for {
		// broadcast로부터 수신한 Message객체를 msg 변수에 할당
		msg := <-broadcast

		// clients map에 담긴 커넥션 객체를 순회
		for conn := range clients {
			// 각 커넥션에 msg 값을 작성해 전달
			err := conn.WriteJSON(msg)
			// 에러 처리: 커넥션에서 에러가 발생하면 커넥션을 종료하고 clients map에서 커넥션을 지움
			// 메시지 핸들러는 계속해서 나머지 커넥션의 메시지에 대응해야 하므로
			// 커넥션 핸들러와 달리 루프를 종료하지는 않는다
			if err != nil {
				log.Println("message: ", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}
