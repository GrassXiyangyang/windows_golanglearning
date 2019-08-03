package main

import (
	"fmt"
	"regexp"
)

var str string
var b bool

func IsNumber(str string) bool {
	b = false
	rege1 := regexp.MustCompile(`(^[1-9][0-9]*)|0(\.[0-9]+)?`)
	rege2 := regexp.MustCompile()
	// rege1 := regexp.MustCompile(`^\d*\.{0,1}\d{0,1}$`)
	//intnum := rege2.FindAllStringIndex(str, -1)
	in := rege1.MatchString(str)
	//float1 := rege1.FindAllStringSubmatch(str, -1)
	fmt.Print(in)
	//fmt.Print(float1)
	return b
}

func main() {
	fmt.Println("plase input you str:")
	fmt.Scan(&str)
	// str = ` sdsdsjd
	// 9023
	// 23232.2323
	// 2323.232.332.3.232.3
	// 23232.eee
	// -23232
	// -32dd.
	// +323
	// `
	IsNumber(str)
	//fmt.Print(IsNumber(str))
}
