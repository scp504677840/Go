package main

import "fmt"

/*
封装：通过方法来实现。
继承：通过匿名字段实现。
多态：通过接口实现。
 */
func main() {

	//写法一：分别初始化
	address := Address{1, "bj"}
	student := Student{11, "tom", 20, 666,"hahaha",address}

	//写法二：简写
	//student := Student{11, "tom", 20, 666,"hahaha", Address{1,"bj"}}

	fmt.Println(student.int)//666
	fmt.Println(student.Str)//hahaha
	fmt.Println(student)        //{11 tom 20 666 hahaha {1 bj}}
	fmt.Printf("%v\n", student) //{11 tom 20 666 hahaha {1 bj}}
	//%+v显示更详细
	fmt.Printf("%+v\n", student) //{id:11 name:tom age:20 int:666 Str:hahaha Address:{id:1 city:bj}}

}

/*
自定义数据类型
 */
type Str string

/*
其中Student里面有一个叫 id 的字段，而匿名字段Address里面也有一个叫 id 的字段；
这就属于同名字段。遇到同名字段采取就近原则，也就是说，当Student和Address里面同时存在相同字段的时候，
给同名字段赋值的时候，先去Student里面去找，如果没有找到的话再去Address里面去找。
 */
type Student struct {
	id   int
	name string
	age  int
	int //基本数据类型的匿名字段
	Str //自定义数据类型的匿名字段
	Address
	//匿名字段：只有类型，没有名字。继承Address所有字段。
}

type Address struct {
	id   int
	city string
}
