package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入你的等级（例如 1,3)")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("you salary is 4000/month")
		//在go语言中默认是不用break跳出循环，如不需要跳出循环，用fallthrough
		//fallthrough
	case 2:
		fmt.Println("you salary is 600/month")
		//fallthrough
	case 3:
		fmt.Println("you salary is 800/month")
	default:
		fmt.Println("you are the boss,you not have the salary,but you can pay someone a salary")
	}
	var sorce int
	fmt.Println("请输入你的分数")
	fmt.Scan(&sorce)
	switch { //可以没有执行条件
	case sorce > 90: //在case中设置判断条件
		fmt.Println("优秀")
	case sorce > 80:
		fmt.Println("良好")
	case sorce > 70:
		fmt.Println("中等")
	case sorce > 60:
		fmt.Println("合格")
	default:
		fmt.Println("不合格")
	}

}
