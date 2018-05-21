package main

import "fmt"

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
iota 可以被用作枚举值：
const (
    a = iota
    b = iota
    c = iota
)
 */
func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	//0 1 2
	const (
		d = iota
		e
		f
	)
	fmt.Println(d, e, f)
	//0 1 2
	const (
		g = iota //0
		h //1
		i //2
		j = "ha" //ha，iota + 1 = 3
		k // ha，注意这里不是4(iota + 1 = 4)
		l = 100 //100，(iota + 1 = 5)
		m //100，注意这里还是延续上一个常量的值，(iota + 1 = 6)
		n //100，同上。(iota + 1 = 7)
		o = iota //8 (iota + 1 = 8)
		p //9 (iota + 1 = 9)
		q //10 (iota + 1 = 10)
	)
	fmt.Println(g, h, i, j, k, l, m, n, o, p, q)
	//          0 1 2 ha ha 100 100 100 8 9 10
}
