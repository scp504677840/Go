package main

import "fmt"

/*
Go 语言范围(Range)
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
 */
func main() {

	nums := []int{11, 12, 13, 14, 15, 16}
	fmt.Println(nums)

	for i := range nums {
		fmt.Println("索引：", i, "值：", nums[i])
	}

	fmt.Println("--------------------------------------------------------------------")

	for i, v := range nums {
		fmt.Println("索引：", i, "值：", v)
	}

	fmt.Println("--------------------------------------------------------------------")

	peoples := map[string]int32{"tom": 10, "jay": 20, "char": 35}
	fmt.Println(peoples)

	for k, v := range peoples {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println("--------------------------------------------------------------------")

	//range也可以用来枚举Unicode字符串。
	//第一个参数是字符的索引，第二个是字符(Unicode的值)本身。
	for i, c := range "go" {
		fmt.Println(i, c, string(c))
	}

}
