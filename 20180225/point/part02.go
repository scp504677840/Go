package main

import "fmt"

/*
什么是指针
一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
 */
func main() {
	//指向整型
	var ip *int
	//指向浮点型
	var fp *float64

	fmt.Println(ip)
	fmt.Println(fp)
}
