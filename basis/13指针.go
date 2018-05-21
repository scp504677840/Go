package main

import "fmt"

/*
指针
指针变量
 */
func main() {

	num := 10
	p := &num
	fmt.Println(p)//0xc420080010
	fmt.Println(*p)//10
	fmt.Printf("num address : %p\n", &num)//num address : 0xc420080010
	fmt.Printf("type: %T\n",p)//type: *int

	var p1 *int
	fmt.Println(p1)//<nil>

}
