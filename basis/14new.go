package main

import "fmt"

func main() {

	a := 10
	var p *int
	p = &a
	fmt.Println(p)

	//动态分配空间
	p = new(int)
	*p = 666
	fmt.Println(*p)

	//自动类型推断
	q := new(int)
	*q = 777
	fmt.Println(*q)

}
