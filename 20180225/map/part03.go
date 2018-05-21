package main

import "fmt"

/*
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
 */
func main() {
	a := make(map[string]string)
	a["one"] = "tom"
	a["two"] = "jay"
	a["three"] = "cher"

	for k, v := range a {
		fmt.Println(k, v)
	}
	//three cher
	//one tom
	//two jay

	delete(a, "one")

	for k, v := range a {
		fmt.Println(k, v)
	}
	//two jay
	//three cher
}
