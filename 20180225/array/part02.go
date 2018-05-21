package main

import "fmt"

/*
Go 语言支持多维数组，以下为常用的多维数组声明方式：
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
 */
func main() {
	var a [2][3]int
	var b = [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(a)
	fmt.Println(b)
}
