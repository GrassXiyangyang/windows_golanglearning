package main

import "fmt"

func text_2(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b

	} else {
		max = b
		min = a

	}
	fmt.Printf("min = %d , max = %d\n", min, max)
	return
}

//程序入口
func main() {

	text_2(4, 8)
}
