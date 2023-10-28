package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}

type Resp struct {
	Num int
}

func main() {
	req := Req{
		Num1: 1,
		Num2: 2,
	}
	var resp Resp
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 同步
	//client.Call("Server.Add", req, &resp)
	ca := client.Go("Server.Add", req, &resp, nil)
	fmt.Println("do many things")
	<-ca.Done
	fmt.Println(resp)
}
