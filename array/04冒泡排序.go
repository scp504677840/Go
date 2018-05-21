package main

import "fmt"

/*
冒泡排序
*/
func main() {

	var nums = [10]int{1, 20, 5, 15, 3, 18, 6, 12, 4, 11}

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				var temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}

}
