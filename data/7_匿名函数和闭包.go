package main

import "fmt"

func main() {

	a := 4
	str := "mike"
	//匿名函数：没有函数名
	f1 := func() {
		fmt.Printf("a= %d, str=%s\n", a, str)
	}
	f1() //调用前面定义的匿名函数

	//无参的匿名函数
	func() {
		fmt.Printf("a= %d, str=%s\n", a, str)
	}() //直接调用匿名函数

	//有参数的匿名函数
	func(a, b int) {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(10, 89)
}
