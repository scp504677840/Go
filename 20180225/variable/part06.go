package main

import (
	"fmt"
	"reflect"
)

/*
vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误
 */
func main() {

	a, b, c := 1, 2.0, true
	fmt.Println(a, b, c, reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
	//1 2 true int float64 bool

}
