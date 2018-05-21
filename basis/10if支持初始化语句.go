package basis

import "fmt"

func main() {

	a := "hi"
	var b string
	fmt.Scan(&b)

	if a == b {
		println("ok")
	} else if b == "c" {
		println("ccc")
	}

	var c int
	fmt.Scan(&c)
	if d := c + 10; d < 10 {
		println("if支持初始化语句")
	}

}
