package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func main() {
	var s Student
	fmt.Println(s)//{0  0}
	s = Student{1, "tom", 22}
	fmt.Println(s)//{1 tom 22}

	//指针变量
	var ps *Student
	//情况一：普通写法，直接取结构体地址
	ps = &Student{2, "jack", 23}
	//情况二：取已存在对象当地址
	//ps = &s
	//情况三：通过new直接得到内存地址
	//ps = new(Student)

	fmt.Println(ps)//&{2 jack 23}

	s.name = "mike"
	s.age = 30
	s.id = 3
	fmt.Println(s)//{3 mike 30}

}
