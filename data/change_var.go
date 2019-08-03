package main

import "fmt"

func main() {
	//var a int = 10
	//var b int = 20
	//var c int = 30
	a, b, c := 10, 20, 30
	a, b, c = c, b, a //交换变量
	//对于printf中的%d表示后面的值填入这里
	fmt.Printf("a = %d ,b = %d ,c = %d", a, b, c)

}
