package main

import "fmt"

/*
Go 语言指向指针的指针
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
 */
func main() {
	var a int = 10
	var ptr *int
	var pptr **int
	ptr = &a
	pptr = &ptr
	fmt.Println(a, &a)//10 0xc42001a0d8
	fmt.Println(ptr, *ptr, &ptr)//0xc42001a0d8 10 0xc42000c028
	fmt.Println(pptr, *pptr, **pptr, &pptr)//0xc42000c028 0xc42001a0d8 10 0xc42000c030
}
