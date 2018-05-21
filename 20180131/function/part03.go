package main

import "fmt"

/*
Go 语言函数引用传递值
引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：
 */
func swap3(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("交换前a的值： %d\n", a)
	fmt.Printf("交换前b的值： %d\n", b)

	swap3(&a, &b)

	fmt.Printf("交换后a的值： %d\n", a)
	fmt.Printf("交换后b的值： %d\n", b)

	//交换前a的值： 100
	//交换前b的值： 200
	//交换后a的值： 200
	//交换后b的值： 100

}
