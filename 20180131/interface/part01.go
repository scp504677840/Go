package main

import "fmt"

/*
Go 语言接口
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

定义接口
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	......
	method_namen [return_type]
}

定义结构体
type struct_name struct {
}

实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
}
......
func (struct_name_variable struct_name) method_namen() [return_type] {
}
 */
func main() {

	var phone Phone

	//phone = new(XiaoMi)
	phone = XiaoMi{name: "XM"}
	phone.call()

	phone = new(OPPO)
	phone.call()

}

type Phone interface {
	call()
}

type XiaoMi struct {
	name string
}

func (mi XiaoMi) call() {
	fmt.Println("用小米手机打电话")
}

type OPPO struct {
}

func (oppo OPPO) call() {
	fmt.Println("用OPPO打电话")
}
