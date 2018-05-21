package function

import "fmt"

/*
defer
1.作用域：函数内部
2.函数结束前调用
*/
func main() {

	//一个defer
	/*fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")*/
	//a
	//c
	//b

	//多个defer执行顺序
	/*fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	defer fmt.Println("d")
	fmt.Println("e")*/
	//先进后出
	//a
	//e
	//d
	//c
	//b

	//特殊情况一：当遇到panic时，panic以下的defer不会被执行
	/*defer fmt.Println("a")
	defer fmt.Println("b")
	test(0)
	defer fmt.Println("c")*/
	//b
	//panic: runtime error: integer divide by zero
	//a

	//特殊情况三：当遇到被defer 修饰的 会发生panic函数时，所有的defer都会被执行
	/*defer fmt.Println("a")
	defer fmt.Println("b")
	defer test(0)
	defer fmt.Println("c")*/
	//c
	//b
	//a
	//panic: runtime error: integer divide by zero

}

func test(a int) {
	fmt.Println(100 / a)
}
