package main

import "fmt"

func main() {
	fmt.Printf("请输入的你的选择的：")

	var s string
	fmt.Scan(&s)
	//if 支持值的一个初始化，判断用;分隔
	// if s:="大哥";s=="大哥"{
	// 	fmt.Printf("you chonse is boss! ")
	// }
	//多个if else的选择模式
	if s == "大哥" {
		fmt.Printf("you chonse is boss! ")

	} else if s == "小弟" {
		fmt.Printf("you chonse is borther")
	} else if s == "surper_man" {
		fmt.Printf("you chonse is a man,who's knickers underwear ")
	}

}
