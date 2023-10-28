package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	UP = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
	}
	conns []*websocket.Conn
)

func handle(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conns = append(conns, conn)
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("你说的是："+string(p)+"吗？"))
		}
		fmt.Println(m, string(p), err)
	}
	defer conn.Close()
	log.Println("service closed")
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8888", nil)
}
