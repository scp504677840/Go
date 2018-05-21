package main

import "fmt"

/*
数组之间的比较
 */
func main() {

	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [3]int{1, 2, 3}
	fmt.Println(a == b) //false
	fmt.Println(a == c) //true

}
