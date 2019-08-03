package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "conncet success !")
	//读取用户数据
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("主机ip[%s]: %s\n", addr, string(buf[:n]))
		fmt.Println(len(string(buf[:n])))
		if string(buf[0:n-2]) == "exit" {
			fmt.Println("the end is:", string(buf[0:n-2]))
			return
		}
		//把接受数据转化为大写返回给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listener.Close()
	//接受多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		go HandleConn(conn)
	}
}
