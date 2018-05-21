package main

import "fmt"

/*
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
 */
/*switch x.(type){
case type:
statement(s);
case type:
statement(s);
*//* 你可以定义任意个数的case *//*
default: *//* 可选 *//*
statement(s);
}*/
func main() {
	var a interface{}
	//a = func(int) float64 {return 1}
	switch i := a.(type) {
	case nil:
		fmt.Printf("a的类型是：%T", i)
	case int:
		fmt.Printf("a的类型是int")
	case float64:
		fmt.Printf("a的类型是float64")
	case func(int) float64:
		fmt.Printf("a的类型是func(int) float64")
	case bool, string:
		fmt.Printf("a的类型是bool或string")
	default:
		fmt.Printf("未知类型")
	}
}
