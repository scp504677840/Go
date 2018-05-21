package main

import "fmt"

/*
Go 语言指向指针的指针
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
pointer(address)->pointer(address)->variable(value)
指向指针的指针变量声明格式如下：
var ptr **int
以上指向指针的指针变量为整型。
访问指向指针的指针变量值需要使用两个 * 号，如下所示：
 */
func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	//指针ptr的地址
	ptr = &a

	//指向指针ptr的地址
	pptr = &ptr

	//获取pptr的值
	fmt.Println(ptr)  //0xc42001a0d8 = a 的变量地址
	fmt.Println(*ptr) //3000 = a 的值

	fmt.Println(pptr)   //0xc42000c028 = ptr 的变量地址
	fmt.Println(*pptr)  //0xc42001a0d8 = a 的变量地址
	fmt.Println(**pptr) //3000 = a 的值

}
