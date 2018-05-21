package main

import "fmt"

/*
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
 */
func main() {

	arrays := [3]int{1, 2, 3}
	for i, v := range arrays {
		fmt.Println(i, v)
	}

	list := []int{11, 22, 33}
	for i, v := range list {
		fmt.Println(i, v)
	}

	stu := map[string]interface{}{"name": "tom", "age": 22}
	for k, v := range stu {
		fmt.Println(k, v)
	}

}
