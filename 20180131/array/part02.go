package main

import "fmt"

/*
Go 语言多维数组
Go 语言支持多维数组，以下为常用的多维数组声明方式：
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

以下实例声明了三维的整型数组：
var nums [2][3][4]int

二维数组
二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：
var arrayName [x][y] variable_type
variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列
 */
func main() {

	var nums [2][2]int32

	//初始化二维数组
	/*nums[0][0] = 0
	nums[0][1] = 1
	nums[1][0] = 9
	nums[1][1] = 10*/

	nums = [2][2]int32{
		{1, 11},
		{2, 22},
	}

	//访问二维数组
	//fmt.Println(nums[0][0])

	//遍历二维数组
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			fmt.Println(nums[i][j])
		}
	}

}
