package main

import "fmt"

/*
二维数组
 */
func main() {

	/*var nums [3][4]int
	for _, arr := range nums {
		for _, v := range arr {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}*/
	//0	0	0	0
	//0	0	0	0
	//0	0	0	0

	var nums [3][4]int = [3][4]int{{1, 2, 3, 4}, {}, {}}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Printf("%d\t", nums[i][j])
		}
		fmt.Println()
	}
	//1	2	3	4
	//0	0	0	0
	//0	0	0	0

}
