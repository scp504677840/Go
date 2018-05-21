package function

import "fmt"

/*
defer 与 匿名函数
有两种情况：
1.无参匿名函数访问外部变量，延迟调用时，读实时内存变量
2.有参匿名函数，传递给匿名函数的实参在程序读到defer 匿名函数时已经确定，只是没有被调用，延迟调用时，直接调用此匿名函数。
*/
func main() {

	a := 11
	b := true

	defer func(a int, b bool) {
		fmt.Printf("defer: a = %d,b = %t\n", a, b)
	}(a, b)
	//这里特别注意：此函数参数在程序读到这里的时候已经被确定了，只是没有被调用。

	a = 22
	b = false
	fmt.Printf("a = %d,b = %t\n", a, b)

	//a = 22,b = false
	//defer: a = 11,b = true

}

func test09() {
	a := 11
	b := true

	defer func() {
		fmt.Printf("defer: a = %d,b = %t\n", a, b)
	}()

	a = 22
	b = false
	fmt.Printf("a = %d,b = %t\n", a, b)

	//a = 22,b = false
	//defer: a = 22,b = false
}
