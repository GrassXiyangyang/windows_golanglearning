package main

import "fmt"

func conumster(ch chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y

		//如果有输入到quit中时，结束
		case <-quit:
			return

		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}
		//结束的时候，判断什么时候出来
		quit <- true

	}()
	conumster(ch, quit)

}
