package main

import (
	"fmt"
	"time"
)

var apples = make(chan int) //存储所有的苹果

/*
放入苹果
 */
func save(apple int) {
	apples <- apple
}

/*
吃苹果
 */
func eat() {
	for {
		select {
		case apple := <-apples:
			fmt.Println(apple)
		}
	}
}

func main() {

	//存入1000个苹果
	for i := 0; i < 1000; i++ {
		go save(i)
	}

	//生成10个人开始吃
	for i := 0; i < 10; i++ {
		go eat()
	}

	<-time.After(5 * time.Second)

}
