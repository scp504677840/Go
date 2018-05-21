package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	//第一步：设置种子

	//固定种子，如果是固定种子的话，每次产生的随机数都是相同的。
	//rand.Seed(666)

	//用当前时间的毫秒值作为种子，每次产生的随机数都是不相同的。
	rand.Seed(time.Now().UnixNano())

	//第二步：产生随机数，伪随机数
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

}
