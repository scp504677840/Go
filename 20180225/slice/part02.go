package main

import "fmt"

func main() {

	a := make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println(a)

	b := a[:]
	fmt.Println(b)

	c := a[1:]
	fmt.Println(c)

	d := a[:9]
	fmt.Println(d)

	e := a[2:5]
	fmt.Println(e)

}
