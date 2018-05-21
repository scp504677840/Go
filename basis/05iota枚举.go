package basis

import "fmt"

/*
iota：常作为枚举使用
1.只能给常量使用
2.自动加1
3.遇到const重置为0
4.同一行，值相同
 */
const (
	A    = iota
	B    = iota
	C    = 10
	D
	E
	F, G = iota, iota
)

func main() {
	fmt.Println("A", A)
	fmt.Println("B", B)
	fmt.Println("C", C)
	fmt.Println("D", D)
	fmt.Println("E", E)
	fmt.Println("F", F)
	fmt.Println("G", G)
	// A 0
	// B 1
	// C 10
	// D 10
	// E 10
	// F 5
	// G 5
}
