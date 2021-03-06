package main

import "fmt"

/*
Go 语言 select 语句 通信操作
select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
以下描述了 select 语句的语法：
#1 每个case都必须是一个通信
#2 所有channel表达式都会被求值
#3 所有被发送的表达式都会被求值
#4 如果任意某个通信可以进行，它就执行；其他被忽略。
#5 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
   否则：
   #1 如果有default子句，则执行该语句。
   #2 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
 */
func main() {

	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Println("从c1接收的值为： ", i1)
	case c2 <- i2:
		fmt.Println("发送i2的值给c2 ", i2)
	case i3, ok := <-c3:
		if ok {
			fmt.Println("从c3接收的值为：", i3)
		} else {
			fmt.Println("c3已经关闭")
		}
	default:
		fmt.Println("其它")
	}

}
