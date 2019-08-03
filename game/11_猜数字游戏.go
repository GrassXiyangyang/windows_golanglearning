package main

import "fmt"
import "math/rand"
import "time"

func CreatNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	*p = rand.Intn(9000) + 1000 //
	//fmt.Println("randNum=", *p)

}
func GetNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10

}

func Game(randsclic []int) {
	var num int
	ScanSclic := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个四位数")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			}
			fmt.Println("请按照要求输入数字！")

		}
		GetNum(ScanSclic, num)
		//fmt.Println("ScanSclic = ", ScanSclic)
		n, m := 0, 0
		for i := 0; i < 4; i++ {

			if ScanSclic[i] > randsclic[i] {

				fmt.Printf("你输入的第[%d]位的数字大了\n", i)

			} else if ScanSclic[i] < randsclic[i] {

				fmt.Printf("你输入的第[%d]位的数字小了\n", i)

			} else if ScanSclic[i] == randsclic[i] {

				fmt.Printf("你输入的第[%d]位的数字正确\n", i)
				n++
			}
			m = m + 1

		}
		if n == 4 {
			fmt.Printf("恭喜你在第%d次猜对了,答案是：", m)

			fmt.Println("randsclic = ", randsclic)
			break
		}

	}

}
func main() {

	var randNum int
	CreatNum(&randNum)
	//fmt.Println("randNum=", randNum)
	//切片保存产生随机数的每一位
	randsclic := make([]int, 4)
	//取出每一位数
	GetNum(randsclic, randNum)
	fmt.Println("randsclic = ", randsclic)
	//输入一个4位数于随机产生的4位数进行比较
	Game(randsclic)

}
