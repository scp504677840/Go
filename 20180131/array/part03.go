package main

import "fmt"

/*
Go 语言向函数传递数组
如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：
#1 形参设定数组大小：
func myFunction(param [10]int)
#2 形参未设定数组大小：
func myFunction(param []int)
 */
func main() {

	var nums1 = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	plOne(nums1)

	fmt.Println("-------------------------------------------------------------------")

	var nums2 = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	plTwo(nums2)

}

func plOne(nums [10]int) {
	for i, v := range nums {
		fmt.Println("索引：", i, "值：", v)
	}
}

func plTwo(nums []int) {
	for i, v := range nums {
		fmt.Println("索引：", i, "值：", v)
	}
}
