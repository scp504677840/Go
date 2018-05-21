package main

import "fmt"

/*
Go 语言指针数组
 */
func main() {

	a := [3]int{1, 2, 3}
	var ptr [MAX]*int

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}

	for k, v := range ptr {
		fmt.Println(k, v, *v)
	}
	//0 0xc420016380 1
	//1 0xc420016388 2
	//2 0xc420016390 3

}

const MAX int = 3
