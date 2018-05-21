package main

import "fmt"

/*
指针传递数组，修改方法内改变数组，外部数组也改变。
 */
func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	swap2(&arr)
	fmt.Println(arr)
	//[1 2 3 4 5]
	//&[55 2 3 4 5]
	//[55 2 3 4 5]
}

func swap2(pArr *[5]int) {
	pArr[0] = 55
	fmt.Println(pArr)
}
