package main

import "fmt"

/*
Go 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr。
 */
func main() {

	var ptr *int
	fmt.Println(ptr)
	fmt.Printf("ptr 的值为： %x\n", ptr)
	fmt.Println("ptr是否为空指针:",ptr==nil)

}
