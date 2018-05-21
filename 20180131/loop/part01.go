package main

import "fmt"

/*
Go 语言 for 循环
for循环是一个循环控制结构，可以执行指定次数的循环。
Go语言的For循环有3中形式，只有其中的一种使用分号。

和 C 语言的 for 一样：
for init; condition; post { }
init： 一般为赋值表达式，给控制变量赋初值；
condition： 关系表达式或逻辑表达式，循环控制条件；
post： 一般为赋值表达式，给控制变量增量或减量。

和 C 的 while 一样：
for condition { }

和 C 的 for(;;) 一样：
for { }

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
for key, value := range oldMap {
		newMap[key] = value
}
 */
func main() {

	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	/*var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}*/

	/*for {
		fmt.Println("无限循环")
	}*/

	/*numbers := [6]int32{1, 2, 3, 4, 5, 6}
	for i, x := range numbers {
		fmt.Printf("索引：%d，值：%d\n", i, x)
	}*/

	//Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
	var i int32 = 0

LOOP:
	for i < 10 {
		i++
		if i == 5 {
			i++
			goto LOOP
		}
		fmt.Println(i)
	}

}
