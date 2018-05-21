package main

import "fmt"

/*
append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
 */
func main() {
	var nums []int
	pl3(nums)

	//允许追加空切片
	nums = append(nums, 1)
	pl3(nums)

	//向切片中追加一个元素
	nums = append(nums, 2)
	pl3(nums)
	nums = append(nums, 3)
	pl3(nums)
	nums = append(nums, 4)
	pl3(nums)

	//同时添加多个元素
	nums = append(nums, 5, 6, 7)
	pl3(nums)

	//创建切片nums2是之前切片的两倍容量
	nums2 := make([]int, len(nums), cap(nums)*2)
	pl3(nums2)

	//拷贝nums的内容到nums2
	copy(nums2, nums)
	pl3(nums2)

}

func pl3(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
}
