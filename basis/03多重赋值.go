package basis

import "fmt"

/*
多重赋值
 */
func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	// 交换
	a, b = b, a
	fmt.Println(a, b)
}
