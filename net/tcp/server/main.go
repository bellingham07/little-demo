package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建tcp连接
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	// 监听端口
	listen, _ := net.ListenTCP("tcp", addr)
	// 开始连接
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 服务端接收信息
		go handleConn(conn)
	}
}

func handleConn(conn *net.TCPConn) {
	for {
		rr := make([]byte, 1024)
		n, err := conn.Read(rr)
		if err != nil {
			return
		}
		fmt.Println(conn.RemoteAddr().String() + "说：" + string(rr[:n]))
		// 向客户端发送信息
		str := "收到了" + string(rr[:n])
		conn.Write([]byte(str))
	}
}
