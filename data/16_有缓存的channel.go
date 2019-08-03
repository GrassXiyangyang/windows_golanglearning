package main

import (
	"fmt"
	//"time"
)

//close channel

func main() {

	var ch = make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i

			//fmt.Printf("ch[%d]=%d\n", i, i)
		}
		close(ch)
	}()
	for data := range ch {
		fmt.Printf("ch[]=%d\n", data)
	}
}

// //close channel

// func main() {

// 	var ch = make(chan int, 3)

// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			ch <- i

// 			//fmt.Printf("ch[%d]=%d\n", i, i)
// 		}
// 		close(ch)
// 	}()
// 	for {

// 		if i, num := <-ch; num == true {
// 			fmt.Println("num =", i)
// 		} else {
// 			break
// 		}

// 	}
// }

//有缓冲的channel
// func main() {

// 	var ch = make(chan int, 3)

// 	go func() {
// 		for i := 0; i < 12; i++ {
// 			ch <- i
// 			fmt.Printf("ch[%d]=%d\n", i, i)
// 		}

// 	}()
// time.Sleep(time.Second)
// 	for i := 0; i < 12; i++ {
// 		num := <-ch
// 		fmt.Println("num=", num)
// 	}
// }

//生产者消费者模型
// func Producer(queue chan<- int) {//chan<- int 为单向channel 只能写

// 	for i := 0; i < 10; i++ {

// 		queue <- i
// 		fmt.Println("Producer:", i)

// 	}

// }

// func Consumer(queue <-chan int) {//<-chan int 为单向channel 只能读

// 	for i := 0; i < 10; i++ {

// 		v := <-queue

// 		fmt.Println("receive:", v)

// 	}

// }

// func main() {

// 	queue := make(chan int, 1)

// 	go Producer(queue)

// 	go Consumer(queue)

// 	time.Sleep(1e9) //主协程延迟，让Producer与Consumer完成

// }
