package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendfile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open error!", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.read err", err1)
			}
			return
		}
		//发送文件
		conn.Write(buf[:n])

	}

}
func main() {
	fmt.Println("请输入问文件名：")
	var path string
	fmt.Scan(&path)
	//获取文件名
	FileInfo, error := os.Stat(path)
	if error != nil {
		fmt.Println(" os.Stat error!", error)
		return
	}
	//主动链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial error!", err)
		return
	}
	defer conn.Close()

	//先给接受方发送文件名
	_, err1 := conn.Write([]byte(FileInfo.Name()))
	if err1 != nil {

		fmt.Println("conn.Write error!", err1)
		return
	}

	//接受对方的回复
	var n int
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read error!", err2)
		return
	}
	if string(buf[:n]) == "ok" {
		//确认发送文件
		sendfile(path, conn)

	}

}
