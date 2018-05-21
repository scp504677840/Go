package main

import "fmt"

/*
数组是值传递
 */
func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	mid(arr)
	fmt.Println(arr)
	//[1 2 3 4 5]
	//[55 2 3 4 5]
	//[1 2 3 4 5]
}

func mid(arr [5]int) {
	arr[0] = 55
	fmt.Println(arr)
}
