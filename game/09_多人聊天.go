package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//定义成员信息的结构体
type Client struct {
	date_channel chan string
	name         string
	addr         string
}

//定义成员数目的map
var online_members map[string]Client

var message = make(chan string)

func MakeMsg(cli Client, str string) (buf string) {

	buf = "[" + cli.addr + "]" + cli.name + " : " + str
	return
}

func Write_channel(cli Client, conn net.Conn) {

	for msg := range cli.date_channel {

		conn.Write([]byte(msg))
	}

}

func HandleConn(conn net.Conn) {

	defer conn.Close()

	//获取用户的IP地址
	Addr := conn.RemoteAddr().String()

	//创建一个结构体,默认姓名为IP地址，然后将新用户加入map中
	cli := Client{make(chan string), Addr, Addr}
	online_members[Addr] = cli
	//给每个用户发送信息
	go Write_channel(cli, conn)
	//告诉所有用户谁上线了
	message <- MakeMsg(cli, "login")

	//成员主动退出
	quit := make(chan bool)
	//成员超时退出
	time_in := make(chan bool)
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				quit <- true
				fmt.Println("用户下线！err =", err)
				return
			}
			data := string(buf[:n-2])
			if len(data) == 3 && data == "who" {
				conn.Write([]byte("user list:"))
				for _, list := range online_members {
					data = list.addr + ":" + list.name
					conn.Write([]byte(data))
				}
			} else if len(data) >= 9 && data[:8] == "modname:" {
				New_name := strings.Split(data, ":")[1]
				cli.name = New_name
				online_members[Addr] = cli
				conn.Write([]byte("change your name success!"))
			} else {
				message <- MakeMsg(cli, data)
			}

			time_in <- true
		}

	}()

	for {

		select {
		case <-quit:
			delete(online_members, cli.addr) //删除当前用户
			message <- MakeMsg(cli, "退出聊天")
			return
		case <-time_in:

		case <-time.After(30 * time.Second):
			delete(online_members, cli.addr) //删除当前用户
			message <- MakeMsg(cli, "超时退出")
			return
		}

	}

}

//检查错误原因
func CheckError(err error) {
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
}

//将消息发送给每个成员的channel
func Send_To_Members() {
	//给map分配空间
	online_members = make(map[string]Client)
	for {
		msg := <-message

		//遍历map，给每个成员的通信channel写入信息
		for _, members_channel := range online_members {
			members_channel.date_channel <- msg
		}
	}
}
func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	CheckError(err)
	defer listener.Close()
	//不停检测和接受用户请求

	go Send_To_Members()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error =", err)
			continue
		}
		//处理用户请求
		go HandleConn(conn)

	}

}
