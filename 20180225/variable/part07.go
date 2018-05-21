package main

import "fmt"

/*
// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
 */

var (
	a int
	b float64
	c bool
)

func main() {

	var (
		d uint
		e float32
		f bool
	)

	fmt.Println(a) //0
	fmt.Println(b) //0
	fmt.Println(c) //false
	fmt.Println(d) //0
	fmt.Println(e) //0
	fmt.Println(f) //false

}
