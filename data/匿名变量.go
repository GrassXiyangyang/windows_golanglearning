package main

import "fmt"

func return_1() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	var d, e, f int
	d, e, f = return_1()
	fmt.Printf("d = %d, e = %d , f = %d \n", d, e, f)
	//使用匿名变量可以丢弃一个值
	_, e, f = return_1()
	fmt.Printf("e = %d , f = %d \n", e, f)
}
