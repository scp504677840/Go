package basis

import "fmt"

func main() {

	// 情况一：声明常量，带类型
	const A int = 10
	fmt.Println(A)

	// 情况二：声明常量并自动推断类型
	const B = 20
	fmt.Println(B)

	// 情况三：声明多个常量
	const (
		C float32 = 3.14
		D         = 40
	)
	fmt.Println(C)
	fmt.Println(D)

}
