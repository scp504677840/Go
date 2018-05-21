package main

import "fmt"

/*
Go 语言函数值传递值
传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
 */
func main() {
	x, y := 1, 2
	swap3(x, y)
	fmt.Println(x, y)
}

func swap3(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}
