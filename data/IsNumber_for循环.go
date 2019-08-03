package main

import "fmt"

var str string //存储输入字符串
var n int      //存储输入字符串的长度
var b bool

func IsNumber(str string) (b bool) {
	intNum := 0     //int 个数
	floatNum := 0   //小数点个数
	var IsChar bool //判断是否含有字母
	IsChar = true   //初始化
	b = false       //初始化最终判断
	for i := 0; i < n; i++ {
		//判断每一位是否有字母，如果有，直接false
		if str[i] >= 'A' && str[i] <= 'z' {
			//如果出现字母为false，为后面的”2389.dshi“情况做判断
			IsChar = false
			//fmt.Println("IsChar =", IsChar)
			break
		}
		//判断每一位是不是数字
		if str[i] >= '0' && str[i] <= '9' {
			//fmt.Println("intNum=", intNum)
			//如果是数字，intNum+1，最后和输入字符长度判断大小，若相等，则每一位都是数字，为int
			intNum = intNum + 1
		} else if str[i] == '.' {
			//判断每一位是不是”.“
			if str[i] == '.' {
				//计算有多少个小数点
				floatNum = floatNum + 1
			}
		} else { //剩下的情况 ：不是字母，不是小数点，不是数字，那就直接为false，例如：+3478，—324,直接break
			b = false
			break
		}
		//此时输入串中只能有数字和".","-","+"

	}
	//判断数是不是int
	if intNum == n {
		fmt.Println("the str type is int! ")
		b = true
	}
	//判断数是不是float
	//先判断数中有没有字母，防止出现”79349.dshjsf“类型也只有一个小数点的情况判断为浮点型
	if IsChar == true {
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
