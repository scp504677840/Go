package main

import "fmt"

/*
Go 语言函数引用传递值
引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递
 */
func main() {
	x, y := 1, 2
	swap4(&x, &y)
	fmt.Println(x, y)
	//2 1
}

func swap4(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
