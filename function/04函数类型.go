package function

import "fmt"


/*
函数类型
1.函数就是一种类型
2.我们可以给函数起别名
3.我们可以声明一个函数类型，然后定义一个这个函数的实现函数，并赋值给它
4.可以实现函数多态
 */
func main() {

	var f = func() {
		fmt.Println(123)
	}
	f()

	//函数类型：多态
	//其中Add函数类型并没有具体函数实现
	//而我们的addPlus函数就是Add函数类型的具体实现
	var plus Add
	plus = addPlus
	fmt.Println(plus(1, 2))

}

type Add func(a, b int) int

func addPlus(a, b int) int {
	return a + b
}
