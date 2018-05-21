package main

import "fmt"

/*
Go 语言向函数传递数组
 */
func main() {
	var a = []int{1, 2, 3, 4, 5}
	//array01(a) //Error
	array02(a)
}

/*
方式一
形参设定数组大小：
 */
func array01(numbers [3]int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}

/*
方式二
形参未设定数组大小：
 */
func array02(numbers []int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}
