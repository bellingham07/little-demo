package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}
	// 客户端发送信息
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		conn.Write(bytes)
		// 接收服务端信息
		buf := make([]byte, 1024)
		conn.Read(buf)
		fmt.Println(string(buf))
	}
}
