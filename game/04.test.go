package main

import "fmt"

func judge(num []int, ch chan int, quit chan bool) {
	x := 0
	fmt.Println("取值为：", num[0])
	for {
		select {
		case y := <-ch:
			if num[x] != y {
				x = x + 1
				num[x] = y
				fmt.Println("取值为：", num[x])
			}
		//如果有输入到quit中时，结束
		case <-quit:
			return
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	num := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 6, 78, 9}
	go func() {
		for i := 0; i < len(num); i++ {
			ch <- num[i]
		}
		//结束的时候，判断什么时候出来
		quit <- true
	}()
	judge(num, ch, quit)

}
