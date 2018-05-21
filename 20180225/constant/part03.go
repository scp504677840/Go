package main

import "fmt"

/*
常量还可以用作枚举：
const (
    Unknown = 0
    Female = 1
    Male = 2
)
 */
func main() {
	const (
		RED   = 0
		GREEN = 1
		BLUE  = 2
	)
	fmt.Println(RED, GREEN, BLUE)
}
