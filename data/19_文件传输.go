package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("the false is 2")
		return
	}
	flieName := list[1]
	FileInfo, error := os.Stat(flieName)
	if error != nil {
		fmt.Println("error!", error)
		return
	}
	fmt.Println("name=", FileInfo.Name())
	fmt.Println("size=", FileInfo.Size())
}
