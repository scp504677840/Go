package main

import "fmt"

/*
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

定义结构体
结构体定义需要使用 type 和 struct 语句。
struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。
type 语句设定了结构体的名称。结构体的格式如下：

type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}

 */
func main() {
	var s1 student
	var s2 student
	fmt.Println(s1)
	fmt.Println(s2)

	s1 = student{"123", "tom", 22}
	fmt.Println(s1)
}

type student struct {
	id   string
	name string
	age  int
}
