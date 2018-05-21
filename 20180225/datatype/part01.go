package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int
	var b int8
	var c int16
	var d int32
	var e int64
	var f uint
	var g uint8
	var h uint16
	var i uint32
	var j uint64
	var k uintptr

	fmt.Println(a, unsafe.Sizeof(a)) //0 8
	fmt.Println(b, unsafe.Sizeof(b)) //0 1
	fmt.Println(c, unsafe.Sizeof(c)) //0 2
	fmt.Println(d, unsafe.Sizeof(d)) //0 4
	fmt.Println(e, unsafe.Sizeof(e)) //0 8
	fmt.Println(f, unsafe.Sizeof(f)) //0 8
	fmt.Println(g, unsafe.Sizeof(g)) //0 1
	fmt.Println(h, unsafe.Sizeof(h)) //0 2
	fmt.Println(i, unsafe.Sizeof(i)) //0 4
	fmt.Println(j, unsafe.Sizeof(j)) //0 8
	fmt.Println(k, unsafe.Sizeof(k)) //0 8

}
