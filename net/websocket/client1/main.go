package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://localhost:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn)
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(m, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, line)
	}
}
