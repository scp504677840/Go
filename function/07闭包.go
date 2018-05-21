package function

import "fmt"

func main() {
	a := 666
	b := "ok"

	func() {
		//闭包以引用方式捕获外部变量
		a = 777
		b = "no"
		fmt.Println("内部：", a, b)
	}()

	fmt.Println("外部：", a, b)

	f := nm()
	fmt.Println(f())//1
	fmt.Println(f())//4
	fmt.Println(f())//9
	fmt.Println(f())//16
	fmt.Println(f())//25
}

/*
演示闭包
闭包：它不关心这些捕获了的变量和常量是否已经超出了作用域，所以只要闭包还在使用它，这些变量就还会存在。
 */
func nm() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
