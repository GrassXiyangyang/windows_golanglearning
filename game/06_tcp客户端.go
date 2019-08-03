package main

import (
	"fmt"
	"net"
)

func main() {
	//主动链接服务器
	conn, error := net.Dial("tcp", "127.0.0.1:8000")
	if error != nil {
		fmt.Println("dial error :", error)
		return
	}
	conn.Write([]byte, "hello")
	//发送请求
	defer conn.Close()
}
