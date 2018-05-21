package main

import "fmt"

/*
iota另外一种用法
 */
func main() {
	const (
		a = iota * 2 // 0 * 2
		b // 1 * 2
		c // 2 * 2
	)
	fmt.Println(a, b, c)
	//0 2 4
}
