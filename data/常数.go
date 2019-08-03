package main

import "fmt"

func main() {

	//常量申明 const
	//变量申明 var
	const a int = 10
	fmt.Println("a=", a)
	var (
		c int     = 45
		b float64 = 8.34
	)
	var e int = 7
	// c, b = 12, 3.98
	fmt.Printf("the type is %T \n", a)
	fmt.Printf("c= %d  ,e= %d\n ", c, e)
	//对于float类型的变量，不可以用
	fmt.Println("b=", b)

}
