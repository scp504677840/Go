package basis

import "fmt"

/*
变量的声明
*/
func main() {

	// 情况一：只声明变量->默认值。
	var a int
	fmt.Println(a)

	// 情况二：声明变量后，赋值。
	var b int
	b = 10
	fmt.Println(b)

	// 情况三：声明并初始化。带类型。
	var c int = 20
	fmt.Println(c)

	// 情况四：声明并初始化。不带类型。
	var d = 30
	fmt.Println(d)

	// 情况五：自动推断类型。
	e := 40
	fmt.Println(e)

}
