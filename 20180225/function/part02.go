package main

import "fmt"

/*
函数返回多个值
 */
func main() {
	x := "a"
	y := "b"
	fmt.Println(swap(x, y))
}

func swap(x, y string) (string, string) {
	return y, x
}
