package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//创建结构体，存储每一个成员信息
type Client struct {
	data_channel chan string //每个用户传输数据管道
	Name         string      //用户名
	Addr         string      //用户网络地址
}

//创建一个map ，保存在线用户
var online_members map[string]Client

//通信管道
var message = make(chan string)

//相关函数定义

//给当前客户端发送信息
func WriteMsgToClient(cli Client, conn net.Conn) {

	for msg := range cli.data_channel {
		//给当前客户端发送信息
		conn.Write([]byte(msg))
	}
}

//传入登陆信息
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return

}

//处理用户链接
func HandleConn(conn net.Conn) { //处理用户链接

	//获取客户端网络地址
	defer conn.Close()
	ClientAddr := conn.RemoteAddr().String()

	//创建一个结构体,默认用户名和网络地址一样
	cli := Client{make(chan string), ClientAddr, ClientAddr}

	//把结构体添加到map
	online_members[ClientAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个在线
	message <- MakeMsg(cli, "已登陆上线")

	cli.data_channel <- MakeMsg(cli, "who l am")

	//成员主动退出
	quit := make(chan bool)
	//成员超时退出
	time_out := make(chan bool)
	//转发成员消息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				quit <- true
				fmt.Println("聊天结束，对方下线 ！err=", err)
				return
			}
			msg := string(buf[:n-2])
			//fmt.Println("服务器返回值长度:", len(msg))
			//用“who”进行当前在线用户查询
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list:"))
				for _, list_member := range online_members {
					msg = list_member.Addr + " : " + list_member.Name + "\n"
					conn.Write([]byte(msg))
				}
				//更改用户名
			} else if len(msg) >= 9 && msg[:8] == "modname:" {
				//modname:yoyo
				new_name := strings.Split(msg, ":")[1]
				cli.Name = new_name
				online_members[ClientAddr] = cli
				//fmt.Println("name : ", cli.Name)
				conn.Write([]byte("modfiy you name success!"))

			} else {
				message <- MakeMsg(cli, msg)
			}
			time_out <- true
		}
	}()

	//不让链接关闭
	for {
		select {
		case <-quit:

			delete(online_members, ClientAddr) //删除当前用户
			message <- MakeMsg(cli, "退出聊天")
			return
		case <-time_out:

		case <-time.After(30 * time.Second):

			message <- MakeMsg(cli, "退出聊天")
			return
		}

	}

}

//广播新消息。有新消息时，把消息转发给每一个成员的信息管道
func forward_message() {

	//给map分配空间
	online_members = make(map[string]Client)
	//一直准备发送消息
	for {
		msg := <-message //没有收到消息前，阻塞,有消息时，读取消息

		//遍历map，将消息传递给map中每一个的channel
		for _, client_members := range online_members {
			client_members.data_channel <- msg
			//fmt.Println("name:", Message_To_Channel.Name)
		}
	}
}

//检查错误信息
func CheckError(err error) {
	if err != nil {
		fmt.Println("error =", err)
		os.Exit(1) //出现错误抛出异常
	}
}

//主函数
func main() {
	//监听
	listener, error := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(error)

	defer listener.Close()

	//新开一个协程，进行广播消息。只要有消息来了，就遍历map，给所有成员发送消息
	go forward_message()

	//阻塞等待用户请求
	//主协程
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error =", err)
			continue
		}

		// 处理用户链接，新来一个用户，开一个协程处理链接问题
		go HandleConn(conn)

	}
}
