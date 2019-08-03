package main

import "fmt"
import "time"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i

	}
	close(out)
}
func conumster(in <-chan int) {
	for data := range in {
		fmt.Println("data =", data)
	}

}
func main() {

	var ch = make(chan int)
	//生产者
	go producer(ch)
	//消费者
	go conumster(ch)
	time.Sleep(1 * time.Minute)
}
