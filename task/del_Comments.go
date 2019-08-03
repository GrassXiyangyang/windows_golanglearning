package main

import (
	"fmt"
	"regexp"
	"strings"
)

func del_SQL_comments(sql string) (str string) {
	rege1 := regexp.MustCompile(`-- \w+`)
	find := rege1.FindAllString(sql, -1)
	//fmt.Println("find = ", find)

	for _, data := range find {
		//fmt.Println(data)
		//用空内容替换前面查找到的注释行，也就是删除注释行
		str = strings.Replace(sql, data, "comments", -1)
		sql = str
	}
	//删除后格式整理
	str = strings.Replace(str, "comments\n", "", -1) //删除后进行格式调整，只去除注释后的换行
	str = strings.Replace(str, "    ", "", -1)       //去除俩个一起的空格
	fmt.Printf(str)
	return
}

func main() {
	sql := `SELECT  emp_no, first_name, last_name 
       FROM employees  inner join cccc on cccc.id = employees.id inner join stusent on cccc.id = student.id 
       -- hahahahahah
    -- hahahahahah
       where last_name='Aamodt' and gender='F' and birth_date > '1960-01-01'   -- bbbbbbbb
    -- cccccccc
    `
	del_SQL_comments(sql)

}
