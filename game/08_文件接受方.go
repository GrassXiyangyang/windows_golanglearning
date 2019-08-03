package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func recevefile(fliename string, conn net.Conn) {
	f, err := os.Create(fliename)
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Write(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接受完毕")
			} else {
				fmt.Println("f.write err", err1)
			}
			return
		}
		fmt.Println("the n =", n)
		if n == 2 {
			fmt.Println("文件接受完毕")
			return
		}
		//发送文件
		conn.Read(buf[:n])

	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	//堵塞等待用户请求
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err =", err1)
		return
	}
	buf := make([]byte, 1024)
	var n int
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err =", err2)
		return
	}
	fliename := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接受文件内容
	recevefile(fliename, conn)
}
