package main

import "fmt"

//变量名 ... 参数类型 为不定参数，根据你传入的参数确定有几个参数
func myfunc1(a ...int) {
	fmt.Printf(" a的长度为：%d\n", len(a))
	//
	for i, data := range a {
		fmt.Printf("a [%d]=%d\n", i, data)
	}
	fmt.Println("------------------")
}

func myfunc2(args ...int) {
	//不定参数全部传递给另一个函数
	myfunc3(args...)
	//不定参数某些值传递给另一个函数
	myfunc3(args[:2]...)
	myfunc3(args[2:]...)
}
func myfunc3(tmp ...int) {
	for _, data := range tmp {
		fmt.Println(" tmp =", data)
	}
	fmt.Println("----------------------")
}

func main() {
	myfunc1()
	myfunc1(0, 1)
	myfunc2(0, 1, 2, 4, 5, 6, 7, 8, 9)
}
