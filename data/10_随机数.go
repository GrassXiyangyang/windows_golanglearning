package main

import "fmt"
import "math/rand"
import "time"

func main() {
	//设置种子

	rand.Seed(time.Now().UnixNano()) //以当前时间为种子，每一次生成的种子都不一样，因为相同种子每次运行产生的随机数相同
	var a [10]int
	n := len(a)

	for i := 0; i < n; i++ {
		// fmt.Println("rand =", rand.Intn(100)) //Intn()中参数设定每次产生随机数的范围
		a[i] = rand.Intn(100)
		fmt.Printf("%d ", a[i])
	}

	fmt.Println("\n排序后结果：------------")

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	// for i := 0; i < n; i++ {

	// 	fmt.Printf("%d ", a[i])
	// }

	fmt.Printf("%d ", a)

}
