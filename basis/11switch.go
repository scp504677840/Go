package basis

import "fmt"

/*
switch：匹配条件，可以有初始化条件。
case：匹配项
break：跳出
fallthrough：继续往下执行
default：没有匹配到匹配项时，执行default
 */
func main() {

	var a int
	fmt.Scan(&a)

	//情况一：正常匹配
	switch a {
	case 1:
		fmt.Println("two")
		//break 默认包含，不需要显式写上
		//break
	case 2:
		fmt.Println("one")
		//匹配多个条件
	case 3, 4:
		fmt.Println("three four")
		//fallthrough不跳出代码块，继续往下执行，不用判断条件
		fallthrough
	case 5, 6:
		fmt.Println("如果执行到了3,4，我们将也会被执行。")
	default:
		fmt.Println("default")
	}

	//情况二：带初始化
	switch b := "hi"; len(b) {
	case 1:
		fmt.Println("长度为1")
	case 2:
		fmt.Println("长度为2")
	default:
		fmt.Println("其他长度")
	}

	//情况三：无需写匹配条件，利用匹配项去判断
	var c string
	fmt.Scan(&c)
	switch {
	case c == "ok":
		fmt.Println("匹配ok")
	case c == "err":
		fmt.Println("匹配err")
	default:
		fmt.Println("匹配default")
	}

}
