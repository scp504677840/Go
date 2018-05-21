package main

import "fmt"

func main() {
	kvs := map[int]string{1: "tom", 2: "jack", 3: "yo"}

	//key不存在时，则添加
	kvs[4] = "kk"

	//key存在时，则修改
	kvs[1] = "cher"
	fmt.Println(kvs)
	//map[1:cher 2:jack 3:yo 4:kk]

	//获取key某个value
	//第一个返回值是key所对应的value
	//第二个返回值是key是否存在
	value, ok := kvs[0]
	fmt.Println("value:", value, "ok:", ok)
	//value:  ok: false

	//遍历map，当写作一个返回值时，这个返回值代表key。
	for key := range kvs {
		fmt.Println(key)
	}
	//1
	//2
	//3
	//4

	//遍历map，当写作两个返回值时，第一个返回值是key，第二个返回值是value。
	for k, v := range kvs {
		fmt.Println(k, v)
	}
	//1 cher
	//2 jack
	//3 yo
	//4 kk

	//根据key值删除
	delete(kvs, 0)
	fmt.Println(kvs)
	//map[1:cher 2:jack 3:yo 4:kk]

	//根据key值删除
	delete(kvs, 1)
	fmt.Println(kvs)
	//map[2:jack 3:yo 4:kk]

	changeMap(kvs)
	fmt.Println(kvs)
	//map[3:change 4:change 2:change]
}

func changeMap(kv map[int]string) {
	for k := range kv {
		kv[k] = "change"
	}
}
