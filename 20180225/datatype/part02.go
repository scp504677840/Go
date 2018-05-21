package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a float32
	var b float64

	//32 位实数和虚数
	var c complex64

	//64 位实数和虚数
	var d complex128

	fmt.Println(a, unsafe.Sizeof(a))//0 4
	fmt.Println(b, unsafe.Sizeof(b))//0 8
	fmt.Println(c, unsafe.Sizeof(c))//(0+0i) 8
	fmt.Println(d, unsafe.Sizeof(d))//(0+0i) 16

}
