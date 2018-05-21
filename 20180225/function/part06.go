package main

import "fmt"

/*
Go 语言函数闭包
Go 语言支持匿名函数，可作为闭包。
匿名函数是一个"内联"语句或表达式。
匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。
该函数的目的是在闭包中递增 i 变量，代码如下：
 */
func main() {
	sequence := getSequence()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
