package main

import "fmt"

/*
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
for key, value := range oldMap {
    newMap[key] = value
}
 */
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	for k, v := range nums {
		fmt.Println(k, v)
	}
}
