package main

import "fmt"

/*
_ 实际上是一个只写变量
 */
func main() {
	_, a := 5, 7
	fmt.Println(a) //7
}
