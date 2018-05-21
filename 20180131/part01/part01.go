package main

import "fmt"

/*
Go 语言条件语句
条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
if 语句由布尔表达式后紧跟一个或多个语句组成。
 */
func main() {

	var a int32 = 30
	if a > 20 {
		fmt.Println("大于")
	}

	fmt.Printf("a 的值为： %d\n",a)

}
