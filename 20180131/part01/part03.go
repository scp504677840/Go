package main

import "fmt"

/*
Go 语言 if 语句嵌套
你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
 */
func main() {

	var a int32 = 100
	var b int32 = 200

	if a < 200 {

		if b < 300 {
			fmt.Println("b小于300")
		} else {
			fmt.Println("b大于300")
		}

		fmt.Println("a小于200")

	}

}
