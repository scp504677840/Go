package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

/*
第二种，根据值自行判定变量类型。
var v_name = value
 */
func main() {

	var a = 2
	fmt.Println(a, unsafe.Sizeof(a), reflect.TypeOf(a))//2 8 int

}
