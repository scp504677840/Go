package main

import "fmt"

/*
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
 */
func main() {
	var a int = 20
	var ip *int
	ip = &a

	fmt.Println(&a)
	fmt.Println(ip)
	fmt.Println(*ip)
}
