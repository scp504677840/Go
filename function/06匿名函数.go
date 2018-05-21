package function

import "fmt"

/*
匿名函数
 */
func main() {

	a := 20
	b := "haha"

	//匿名函数
	f := func() {
		fmt.Println("hi", a, b)
	}
	//手动调用匿名函数
	f()

	//匿名函数自动调用
	func() {
		fmt.Println("hihihi")
	}()

	//匿名函数有形参
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20)

	//匿名函数有形参，有返回值
	var sum = func(a, b int) (sum int) {
		return a + b
	}(10, 20)
	fmt.Println(sum)

	//函数类型 + 匿名函数
	type FuncNM func()
	var f2 FuncNM
	f2 = func() {
		fmt.Println("我是匿名函数")
	}
	f2()

}
