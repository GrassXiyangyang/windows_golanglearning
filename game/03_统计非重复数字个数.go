package main

import "fmt"

func judge(num []int) (sum int) {
	if len(num) == 0 {
		return
	}
	x := 0
	for i := 0; i < len(num); i++ {
		if num[x] != num[i] {
			num[x] = num[i]
			sum++
		}
	}
	fmt.Println("sum =", sum)
	return sum + 1
}
func main() {
	num := []int{}

	fmt.Print(judge(num))
}
