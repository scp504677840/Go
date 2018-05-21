package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

/*
第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
v_name := value
 */
func main() {

	a := 10
	fmt.Println(a, unsafe.Sizeof(a), reflect.TypeOf(a))//10 8 int

}
