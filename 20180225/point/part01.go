package main

import "fmt"

/*
变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
 */
func main() {
	var a int = 10
	fmt.Println(&a)
}
