package main

import "fmt"

func main() {
	//iota是对每一行的常量自动加一赋值。
	//对于同一行的常量，值是一样的，也就是说是按照行数进行加一。
	const (
		a    = iota
		b, e = iota, iota
		c    = iota
	)
	fmt.Printf("a = %d, b = %d, c = %d,e = %d", a, b, c, e)
}
