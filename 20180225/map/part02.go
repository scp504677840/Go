package main

import "fmt"

func main() {

	a := make(map[string]string)
	a["one"] = "tom"
	a["two"] = "jay"
	a["three"] = "cher"

	for key := range a {
		fmt.Println(key)
	}

	for k, v := range a {
		fmt.Println(k, v)
	}

	//其中ok表示元素在集合中是否存在
	v, ok := a["one"]
	fmt.Println(v, ok)
}
