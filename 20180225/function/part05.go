package main

import (
	"math"
	"fmt"
)

/*
Go 语言函数作为值
Go 语言可以很灵活的创建函数，并作为值使用。
以下实例中我们在定义的函数中初始化一个变量，
该函数仅仅是为了使用内置函数 math.sqrt() ，实例为：
 */
func main() {
	lab := func(x float64) float64 {return math.Sqrt(x)}

	//使用函数
	fmt.Println(lab(25))
	//5
}
