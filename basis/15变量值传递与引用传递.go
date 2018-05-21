package main

import "fmt"

func main() {
	//值传递，ab交换不可行
	/*a, b := 10, 20
	fmt.Println(a, b)
	swap1(a, b)
	fmt.Println(a, b)*/

	//指针传递，ab交换可行
	/*a, b := 10, 20
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)*/

	//指针传递，ab交换不可行，因为内部ab交换，外部ab依然无效
	a, b := 10, 20
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	swap3(&a, &b)
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	//0xc42001a0d8 0xc42001a0f0
	//10 20
	//指针交换内部： 0xc42001a0f0 0xc42001a0d8
	//0xc42001a0d8 0xc42001a0f0
	//10 20

}

func swap1(a, b int) {
	a, b = b, a
	fmt.Println("值传递内部：", a, b)
}

func swap2(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("指针传递内部：", *a, *b)
}

func swap3(a, b *int) {
	a, b = b, a
	fmt.Println("指针交换内部：", a, b)
}
