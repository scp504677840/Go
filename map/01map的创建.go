package main

import "fmt"

/*
map的创建
 */
func main() {

	var kvs1 map[string]string
	fmt.Println(kvs1)

	kvs2 := map[string]string{}
	fmt.Println(kvs2, len(kvs2))

	//这里的第二个参数，可以指定长度，只是指定了容量，但是里面却是一个数据也没有。
	kvs3 := make(map[string]string, 10)
	fmt.Println(kvs3, len(kvs3))
	kvs3["name1"] = "tom"
	kvs3["name2"] = "jack"
	kvs3["name3"] = "cher"
	fmt.Println(kvs3, len(kvs3))

	kvs4 := map[int]string{1: "tom", 2: "jack", 3: "cher"}
	fmt.Println(kvs4)

}
