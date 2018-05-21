package main

import "fmt"

/*
数组：
1.定义数组：[n]数据类型
2.n一定要是整型
3.必须指定数组大小
 */
func main() {

	//情况一：声明数组
	var nums [5]int
	for i := 0; i < len(nums); i++ {
		fmt.Println("num:", nums[i])
	}
	//num: 0
	//num: 0
	//num: 0
	//num: 0
	//num: 0

	//情况二：完整初始化
	/*var nums [5]int = [5]int{1,2,3,4,5}
	for k, v := range nums {
		fmt.Println(k,v)
	}*/
	//0 1
	//1 2
	//2 3
	//3 4
	//4 5

	//情况三：简写完整初始化
	/*var nums = [5]int{1, 2, 3, 4, 5}
	for k, v := range nums {
		fmt.Println(k, v)
	}*/
	//0 1
	//1 2
	//2 3
	//3 4
	//4 5

	//情况四：不完整初始化，默认值
	/*var nums = [5]int{1, 2, 3}
	for k, v := range nums {
		fmt.Println(k, v)
	}*/
	//0 1
	//1 2
	//2 3
	//3 0
	//4 0

	//情况五：不完整初始化，指定初始化
	/*var nums = [5]int{2: 20, 4: 40}
	for k, v := range nums {
		fmt.Println(k, v)
	}*/
	//0 0
	//1 0
	//2 20
	//3 0
	//4 40

}
