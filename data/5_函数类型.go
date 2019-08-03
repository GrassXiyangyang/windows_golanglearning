package main

import "fmt"

//函数也是一种数据类型，可以用type给一个函数类型起名
//myfunc_type是一个函数类型
type myfunc_type func(int, int) int

func add(a, b int) int {
	return a + b
}

func main() {
	var result int
	//声明一个函数类型的变量
	//将函数也当作一个数据类型，然后直接输出该类型的结果
	var ftest myfunc_type
	ftest = add
	result = ftest(10, 66)
	fmt.Printf("result = %d", result)
}
