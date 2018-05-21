package function

import "fmt"

/*
函数回调
多态函数
抽象函数
 */
func main() {

	//声明一个函数类型
	var f FuncAdd

	//加法
	f = impl1
	fmt.Println(f(8, 2))
	//10

	//减法
	f = impl2
	fmt.Println(f(8, 2))
	//6

	//乘法
	f = impl3
	fmt.Println(f(8, 2))
	//16

	//除法
	f = impl4
	fmt.Println(f(8, 2))
	//4

}

/*
函数类型：
接收两个int类型，返回一个int类型
这个函数并没有实现，可以视作一个抽象函数
 */
type FuncAdd func(x, y int) int

func impl1(x, y int) int {
	return x + y
}

func impl2(x, y int) int {
	return x - y
}

func impl3(x, y int) int {
	return x * y
}

func impl4(x, y int) int {
	return x / y
}
