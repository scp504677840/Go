package main

import "fmt"

/*
结构体作为函数参数
 */
func main() {
	plStu(SPoint2{"Jay"})
}

type SPoint2 struct {
	name string
}

func plStu(stu SPoint2) {
	fmt.Println(stu)
}
