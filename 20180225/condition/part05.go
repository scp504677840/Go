package main

import "fmt"

/*switch var1 {
case val1:
...
case val2:
...
default:
...
}*/
func main() {
	var a int = 3
	switch a {
	case 1:
		fmt.Println(1)
	case 2, 3:
		fmt.Println(23)
	default:
		fmt.Println(-1)
	}
}
