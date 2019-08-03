package main

import "fmt"

//定义一个函数类型的数据类型
type myfunc_type func(int, int) int

//该类型的函数
func add(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}

//多态函数：调用同一个接口，可以是实现相同类型的具体不同功能
//优点，灵活调用，可以动态增加相同类型的具体功能，也可以先有想法，后面再具体实施方案
func calc(a, b int, ftest myfunc_type) (result int) {
	result = ftest(a, b)
	return
}

func main() {
	a := calc(2, 4, add)
	b := calc(2, 4, subtraction)
	c := calc(2, 4, multiplication)
	d := calc(16, 4, division)
	fmt.Printf("a= %d\n", a)
	fmt.Printf("b= %d\n", b)
	fmt.Printf("c= %d\n", c)
	fmt.Printf("d= %d\n", d)
}
