package main

import "fmt"

/*
结构体指针
你可以定义指向结构体的指针类似于其他指针变量，格式如下：
var struct_pointer *Books
 */
func main() {
	var ptr *SPoint3
	var a SPoint3
	a = SPoint3{"tom"}
	ptr = &a
	fmt.Println(a, ptr)
	fmt.Println(ptr.name)
}

type SPoint3 struct {
	name string
}
