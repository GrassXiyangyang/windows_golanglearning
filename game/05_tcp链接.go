package main

import (
	"fmt"
	"net"
)

func main() {
	//监听

	Listener, error := net.Listen("tcp", "127.0.0.1:8000")
	if error != nil {
		fmt.Println("Listen error = ", error)
		return
	}
	defer Listener.Close()

	//阻塞等待用户链接
	for {

		Conn, error1 := Listener.Accept()
		if error1 != nil {
			fmt.Println("Accept error = ", error1)
			return
		}
		buf := make([]byte, 1024)
		n, err := Conn.Read(buf)
		if err != nil {
			fmt.Println(" Read error = ", err)
			return
		}
		fmt.Println("the buf =", string(buf[:n]))
		defer Conn.Close()
	}

	//接受用户请求

	//

}
