package basis

import "fmt"

/*
类型别名
 */
func main() {

	type bigint int64
	var a bigint
	a = 100
	fmt.Println(a)

	type (
		bigfloat float64
		char byte
	)

	var b bigfloat = 3.14
	fmt.Println(b)

	var c char = 'a'
	fmt.Printf("%c\n", c)

}
