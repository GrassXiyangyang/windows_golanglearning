package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func print(str string) {

	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println(" ")
}
func person1() {
	print("hello,world")
	ch <- 9
}
func person2() {
	<-ch
	print("2222222")
}
func main() {

	go person1()
	go person2()
	//死循环
	for {

	}
}
