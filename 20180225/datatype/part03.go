package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//类似 uint8
	var a byte
	//类似 int32
	var b rune

	fmt.Println(a, unsafe.Sizeof(a))//0 1
	fmt.Println(b, unsafe.Sizeof(b))//0 4
}
