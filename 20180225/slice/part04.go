package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	pl4(nums)
	//len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	//取1:4
	nums1 := nums[1:4]
	pl4(nums1)
	//len=3 cap=8 slice=[1 2 3]
	fmt.Println(nums[:3])

	//取4:
	nums2 := nums[4:]
	pl4(nums2)
	//len=5 cap=5 slice=[4 5 6 7 8]

	nums3 := make([]int,0,5)
	pl4(nums3)
	//len=0 cap=5 slice=[]

	//取:2
	nums4 := nums[:2]
	pl4(nums4)
	//len=2 cap=9 slice=[0 1]

	// 我们可以看出切片，实际的是获取数组的某一部分，
	// len切片<=cap切片<=len数组，切片由三部分组成：指向底层数组的指针、len、cap。
}

func pl4(nums []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)
}
