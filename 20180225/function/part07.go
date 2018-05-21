package main

import "fmt"

/*
带参数的闭包函数调用
 */
func main() {
	add := add7(1, 2)
	fmt.Println(add(10,10))
	fmt.Println(add(10,10))
	fmt.Println(add(10,10))
	//100 3
	//200 3
	//300 3
}

func add7(a, b int) func(x, y int) (int, int) {
	i := 0
	return func(x, y int) (int, int) {
		i++
		return i * x * y, a + b
	}
}
