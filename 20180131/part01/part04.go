package main

import "fmt"

/*
Go 语言 switch 语句
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break
 */
func main() {

	var grade string = "B"
	var marks int32 = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("及格")
	case grade == "D":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}

}
