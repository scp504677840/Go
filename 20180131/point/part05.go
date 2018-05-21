package main

import "fmt"

/*
Go 语言指针作为函数参数
Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。
以下实例演示了如何向函数传递指针，并在函数调用后修改函数内的值，：
 */
func main() {

	var x, y = 10, 20

	fmt.Println("交换前：", x, y)

	swap(&x, &y)

	fmt.Println("交换后：", x, y)

	//交换前： 10 20
	//交换后： 20 10

}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
