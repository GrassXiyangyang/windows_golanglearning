package main

import "fmt"

func main() {
	// a, c := 1, 2
	// defer func() {
	// 	fmt.Printf("内部：a=%d,c=%d\n", a, c)
	// }() //()表示调用匿名函数
	// a, c = 7, 8
	// fmt.Printf("外部：a=%d,c=%d\n", a, c)
	// fmt.Printf("---------------------\n")

	m, n := 1, 2
	defer func(m, n int) {
		fmt.Printf("内部：a=%d,c=%d\n", m, n)
	}(m, n) //此时已经先把参数传入进来了
	m, n = 3, 4
	fmt.Printf("外部：a=%d,c=%d\n", m, n)

}
