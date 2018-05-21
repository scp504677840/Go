package main

import "fmt"

/*
Type Switch
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
Type Switch 语法格式如下：
switch x.(type){
case type:
	statement(s)
case type:
	statement(s)
//你可以定义任意个数的case
default: //可选
	statement(s)
}
 */
func main() {

	var x interface{}
	x = float32(1.23)

	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是： %T\n", i)
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case bool, string:
		fmt.Println("bool string")
	default:
		fmt.Println("其它")
	}

}
