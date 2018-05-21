package main

import "fmt"

/*
Go 语言 if...else 语句
if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
 */
func main() {

	var a int32 = 20

	if a > 30 {
		fmt.Println("大于")
	} else {
		fmt.Println("小于")
	}

	fmt.Printf("a 的值为： %d\n",a)

}
