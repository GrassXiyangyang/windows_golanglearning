package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动链接服务器
	//conn, err := net.Dial("tcp", "127.0.0.1:8000")

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			//把输入内容输入给服务器
			conn.Write(str[:n])

		}

	}()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//fmt.Println("服务器返回值:", string(buf[:n]))
		fmt.Println(string(buf[:n]))
		if string(buf[0:n-2]) == "exit" {
			fmt.Println("the end is:", string(buf[0:n-2]))
			return
		}
	}

}
