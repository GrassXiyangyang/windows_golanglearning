package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Printf("sum = %d \n", sum)
	fmt.Printf("sum = %d\n", sum)
	//range元素的迭代，输出字符串中每一个元素
	//range有俩个返回参数，第一是下标，第二个是元素本身
	str := "abcd"
	for i, data := range str {
		fmt.Printf("第 %d个元素是：%c\n", i, data)
	}
}
