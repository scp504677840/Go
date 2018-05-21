package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	//go 1.9版本对于数字类型，无需定义int及float32、float64，系统会自动识别。
	var a = 1.5
	var b = 2
	fmt.Println(a, unsafe.Sizeof(a), reflect.TypeOf(a))//1.5 8 float64
	fmt.Println(b, unsafe.Sizeof(b), reflect.TypeOf(b))//2 8 int
}
