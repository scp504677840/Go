package main

import "fmt"

/*
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
你可以声明一个未指定大小的数组来定义切片：
var identifier []type

切片不需要说明长度。
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

 */
func main() {

	var a []int
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a == nil)

	var b []int = make([]int, 10)
	fmt.Println(b, len(b), cap(b))

	c := make([]int, 10)
	fmt.Println(c, len(c), cap(c))

	d := make([]int, 10, 20)
	fmt.Println(d, len(d), cap(d))

}
