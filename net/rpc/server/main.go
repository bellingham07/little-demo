package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Req struct {
	Num1 int
	Num2 int
}

type Resp struct {
	Num int
}

func (s *Server) Add(req Req, resp *Resp) error {
	time.Sleep(3 * time.Second)
	resp.Num = req.Num2 + req.Num1
	return nil
}

func main() {
	err := rpc.Register(new(Server))
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
