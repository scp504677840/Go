package basis

import "fmt"

/*
类型转换：只有兼容类型才可以类型转换。
 */
func main() {

	var a = 'a'
	fmt.Println(a)

	var b int
	b = int(a)
	fmt.Println(b)

}
