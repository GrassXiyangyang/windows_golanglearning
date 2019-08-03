package main

import "fmt"
import "time"

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Second)

		if i == 8 {
			fmt.Printf("已经循环了八次，开始暂停！")
			break
		}
		fmt.Printf("i=%d \n", i)

	}
}
