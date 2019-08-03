package main

import "fmt"

func recursion(a int) (result int) {
	if a == 1 {
		return 1
	}
	return a + recursion(a-1)
}

func main() {
	var sum int

	sum = recursion(100)
	fmt.Printf("sum = %d", sum)
}
