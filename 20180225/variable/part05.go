package main

import (
	"fmt"
	"reflect"
)

/*
var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显示声明类型，自动推断
 */
func main() {
	var a, b, c = 1, 2.0, true
	fmt.Println(a, b, c, reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
	//1 2 true int float64 bool
}
