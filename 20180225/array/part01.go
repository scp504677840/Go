package main

import "fmt"

/*
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
var variable_name [SIZE] variable_type
 */
func main() {

	var a [3]int
	fmt.Println(a)

	var b = [3]int{1, 2, 3}
	fmt.Println(b)

	//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
	var c = [...]int{1, 2, 3}
	fmt.Println(c)

}
