package main

import "fmt"
import "time"
import "math/rand"

func CreatNum(p *int) {

	rand.Seed(time.Now().UnixNano())
	*p = rand.Intn(9000) + 1000
}
func GetNum(NUM int, s []int) {

	s[0] = NUM / 1000
	s[1] = NUM % 1000 / 100
	s[2] = NUM % 100 / 10
	s[3] = NUM % 10

}

func Game(randSclic []int) {

	ScaNSclic := make([]int, 4)
	for {
		for {
			var num int
			fmt.Printf("请输入一个4位数:")
			fmt.Scan(&num)
			if num > 999 && num < 10000 { //符合要求的数
				GetNum(num, ScaNSclic)
				fmt.Println("ScanNum = ", num)
				break
			}
			fmt.Printf("请输入一个符合要求的数!\n")
		}
		n := 0
		m := 0
		for i := 0; i < 4; i++ {

			if ScaNSclic[i] > randSclic[i] {
				fmt.Printf("第[%d]位大了\n", i+1)
			} else if ScaNSclic[i] < randSclic[i] {
				fmt.Printf("第[%d]位小了\n", i+1)
			} else if ScaNSclic[i] == randSclic[i] {
				fmt.Printf("第[%d]位猜对了\n", i+1)
				n++
			}
			m = m + 1

		}
		if n == 4 {
			fmt.Printf("恭喜你在第[%d]次猜对了，答案是：", m)
			break
		}
	}

}
func main() {
	var randNum int
	CreatNum(&randNum)
	//fmt.Printf("randNum= %d", randNum)
	randSclic := make([]int, 4) //切片函数处理会直接改变原切片的值
	GetNum(randNum, randSclic)
	fmt.Println("randSclic =", randSclic)
	Game(randSclic)

}
