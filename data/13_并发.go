package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(16) //设定cpu核数
	fmt.Println("n = ", n)
	for {
		go func() {
			fmt.Print(1)

		}()
		fmt.Print(2)
	}

}
