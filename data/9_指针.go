package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("内部： a=%d ,b= %d\n", a, b)
}

func swap_1(a1, b1 *int) {
	*a1, *b1 = *b1, *a1
	fmt.Printf("指针交换： a1=%d ,b1= %d\n", *a1, *b1)
}
func main() {
	a, b := 10, 20

	fmt.Printf("外部原本值： a=%d ,b= %d\n", a, b)
	//值传递，不改变外部变量的值
	swap(a, b)
	fmt.Printf("值交换外部： a=%d ,b= %d\n", a, b)

	//指针的地址传递
	swap_1(&a, &b)
	fmt.Printf("指针外部： a=%d ,b= %d\n", a, b)

}
