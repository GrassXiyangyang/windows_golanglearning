package main

import (
	"fmt"
	"regexp"
)

var str string
var b bool

func IsNumber(str string) bool {
	b = false
	rege1 := regexp.MustCompile(`[0-9]+\.[0-9]+`) //一个小数点
	//rege2 := regexp.MustCompile(`[0-9]+`)         //数字
	rege3 := regexp.MustCompile(`[A-Za-z]+`)     //有字母
	rege4 := regexp.MustCompile(`\w+\.\w+\.\w+`) //多个小数点
	rege5 := regexp.MustCompile(`[^[\.]]|\+|\-`) //判断输入串首位是否为"." 或者是否含有"+"  "-"
	//按照编码进行匹配
	float1 := rege1.FindAllStringSubmatch(str, -1)
	// intnum := rege2.FindAllStringSubmatch(str, -1)
	// intnum := rege2.FindAllIndex(str, -1)
	//fmt.Printf()
	stringnum := rege3.FindAllStringSubmatch(str, -1)
	float2 := rege4.FindAllStringSubmatch(str, -1)
	err1 := rege5.FindAllStringSubmatch(str, -1)
	fmt.Println("err1=", err1)
	if err1 != nil { //如果输入串首位不为"."  并且不含有"+"  "-"，进行一下部判断
		if stringnum == nil { //如果没有字母，进行下一步判断
			if float1 != nil { //判断是否有小数点
				if float2 == nil { //如果有小数点，进一步判断是否有多个小数点
					b = true
					//fmt.Println("float1 = ", float1)
					//fmt.Println("float2 = ", float2)
				}

			} else { //没有字母和小数点，只能是只含有数字
				b = true
				//fmt.Println("intnum = ", intnum)
			}
			//fmt.Print(b)
		} else { //如果有字母，直接false
			b = false
		}
	} else { //输入串首位有"."  或含有"+"  "-"，直接false
		b = false
	}
	return b
}

func main() {
	fmt.Println("plase input you str:")
	fmt.Scan(&str)
	IsNumber(str)
	fmt.Print(IsNumber(str))
}
