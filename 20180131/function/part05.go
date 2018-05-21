package main

import "fmt"

/*
Go 语言函数闭包
Go 语言支持匿名函数，可作为闭包。
匿名函数是一个"内联"语句或表达式。
匿名函数的优越性在于可以直接使用函数内的变量，不必声明。
 */
func main() {

	//nextNumber作为一个函数，函数i为0
	nextNumber := getSequence()
	//fmt.Printf("%T\n", nextNumber)
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2
	fmt.Println(nextNumber1()) //3

}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
