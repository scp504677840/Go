package main

import "fmt"

/*
Go 语言指针作为函数参数
 */
func main() {
	x, y := 1, 2
	point1(&x, &y)
	fmt.Println(x, y)
}

func point1(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
