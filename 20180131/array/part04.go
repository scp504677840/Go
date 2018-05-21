package main

import "fmt"

func main() {

	var nums = []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	plArray(nums)

	changeArray(nums)
	plArray(nums)

}

func plArray(nums []int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func changeArray(nums []int) {
	for i := range nums {
		nums[i] = 99
	}
}
