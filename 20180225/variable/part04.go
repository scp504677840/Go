package main

import "fmt"

/*
多变量声明

//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3
 */
func main() {

	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

}
