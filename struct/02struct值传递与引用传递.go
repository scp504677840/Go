package main

import "fmt"

type Worker struct {
	id   int
	name string
	age  int
}

/*
结构体
判断相等与赋值
值传递
引用传递
 */
func main() {

	var w1 Worker
	var w2 Worker
	fmt.Println(w1 == w2) //true

	w1.id = 1
	w1.name = "tom"
	w1.age = 20

	w2.id = 1
	w2.name = "tom"
	w2.age = 20

	fmt.Println(w1 == w2) //true

	changeWorker(w1)
	fmt.Println(w1) //{1 tom 20}

	changePWorker(&w1)
	fmt.Println(w1) //{1 change point 20}

}

/*
值传递
 */
func changeWorker(worker Worker) {
	worker.name = "change"
}

/*
引用传递
 */
func changePWorker(worker *Worker) {
	worker.name = "change point"
}
