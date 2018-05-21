package main

import "fmt"

/*
Go 语言函数方法
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
		函数体
}
 */
func main() {

	var c Circle
	c.radius = 10
	fmt.Printf("圆形的面积为： %.2f\n", c.getArea())

}

type Circle struct {
	radius float64
}

/*
该method属于Circle类型对象中的方法
 */
func (c Circle) getArea() float64 {
	//c.radius即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}
