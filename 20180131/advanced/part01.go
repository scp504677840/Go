package main

import "fmt"

/*
Go 语言递归函数
递归，就是在运行的过程中调用自己。
 */
func main() {

	fmt.Println(Factorial(uint64(3)))

	fmt.Println(fibonacci(5))

}

/*
阶乘
 */
func Factorial(num uint64) (result uint64) {
	if num > 0 {
		result = num * Factorial(num-1)
		return result
	}
	return 1
}

/*
斐波那契数列
 */
func fibonacci(num int) int {
	if num < 2 {
		return num
	}
	return fibonacci(num-2) + fibonacci(num-1)
}
