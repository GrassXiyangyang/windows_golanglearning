package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		defer fmt.Println("子协程执行结束！")
		for i := 0; i < 4; i++ {
			fmt.Printf("子协程中：i=%d\n", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，我执行结束了"
	}()
	//通过channel来实现主协程合子协程的之间的协调和同步
	end := <-ch
	fmt.Println("子协程传递信息：", end)
}
