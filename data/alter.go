package main

import (
	"fmt"
	"regexp"
	"strings"
)

//var str_1 []string

func IsNumber(str string) (str2 []string) {

	rege1 := regexp.MustCompile(`(?s:(alter table t_data_info))`)
	// rege2 := regexp.MustCompile(`(?s:add column .*)`)
	// rege3 := regexp.MustCompile(`(?s:modify column .*)`)
	// rege4 := regexp.MustCompile(`(?s:add index .*)`)
	// rege5 := regexp.MustCompile(`(?s:drop column .*)`)

	find1 := rege1.FindAllString(str, 1)
	// result3 := rege3.FindAllString(str, -1)
	// result2 := rege2.FindAllString(str, -1)
	// result4 := rege4.FindAllString(str, -1)
	// result5 := rege5.FindAllString(str, -1)
	fmt.Println("find1 = ", find1)
	//fmt.Printf("find1 is a %T", find1)
	for _, data1 := range find1 {
		fmt.Printf("result=%v\n", data1)
	}
	// for _, data2 := range result2 {
	// 	fmt.Printf("result=%v\n", data2)
	// }

	// for _, data3 := range result3 {
	// 	fmt.Printf("result=%v\n", data3)
	// }
	// for _, data4 := range result4 {
	// 	fmt.Printf("result=%v\n", data4)
	// }
	// for _, data5 := range result5 {
	// 	fmt.Printf("result=%v\n", data5)
	// }
	//分割
	s := strings.SplitAfter(str, ";")
	fmt.Println(s)
	// for i, data2 := range s {
	// 	fmt.Println("data2=", data2)
	// 	fmt.Printf("fenge[%d]=%s \n", i, data2)
	// }

	result := strings.Join(s[0:], "  \n"+find1[0])
	fmt.Println("result = ", result)

	return
}

func main() {
	// fmt.Println("plase input you str:")
	// fmt.Scan(&str)
	str := "alter table t_data_info add column xxxx; modify column xxxxx; add index xxxx dddd; drop column haha"
	IsNumber(str)
	//fmt.Print(IsNumber(str))
}
