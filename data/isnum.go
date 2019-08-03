package main

import "fmt"

var str string //存储输入字符串
var n int      //存储输入字符串的长度
var b bool

func IsNumber(str string) (b bool) {
	intNum := 0     //int 个数
	floatNum := 0   //小数点个数
	var ischar bool //判断是否含有字母
	b = false       //初始化最终判断
	for j := 0; j < n; j++ {
		if str[j] >= 'A' && str[j] <= 'z' {
			ischar = false
			break //先判断是否有字母，如果有，直接false
		} else {
			ischar = true
		}
	}
	//此时输入串中只能有数字和".","-","+"
	if ischar == true {
		if str[0] != '.' { //首位不可以为"."
			for i := 0; i < n; i++ {
				if str[i] == '.' { //先判断是否有字母，如果有，直接false
					if str[i] == '.' { //判断有多少个小数点
						floatNum = floatNum + 1
					}
				}
				if str[i] >= '0' && str[i] <= '9' {
					//fmt.Println("intNum=", intNum)
					intNum = intNum + 1
				}

			}
		} else { //剩下的情况都是为false，例如：+3478，—324
			b = false
		}
		if intNum == n {
			fmt.Println("the str type is int! ")
			b = true
		}
		if floatNum == 1 {
			fmt.Println("the str type is float! ")
			b = true
		} else if floatNum >= 2 {
			fmt.Println("错误！有多个小数点！")
		}
	}
	fmt.Print(b)
	return b

}

func main() {
	fmt.Println("plase input you str:")
	fmt.Scan(&str) //输入字符串
	n = len(str)   //获得字符串的长度
	IsNumber(str)  //判断是不是数字类型

}
