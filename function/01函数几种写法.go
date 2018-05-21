package function

import "fmt"

func main() {

	//无参无返回值
	sum1()

	//有参无返回值
	sum2(123)

	//无参有返回值
	fmt.Println(sum3())

	//有参有返回值
	fmt.Println(sum4(2))

	//有多个形参，无返回值
	sum5(1, 2)

	//有多个形参，有返回值
	fmt.Println(sum6(1, 2, true))

	//不定长参数
	fmt.Println(sum7(1, 2, 3))
}

func sum1() {
	fmt.Println("demo")
}

func sum2(a int) {
	fmt.Println(a)
}

func sum3() int {
	return 456
}

func sum4(a int) bool {
	return a > 1
}

func sum5(a, b int) {
	fmt.Println(a, b)
}

func sum6(a, b int, c bool) bool {
	return a > b && c
}

func sum7(args ...int) int {
	fmt.Printf("len:%d\n", len(args))
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

/*
不定长参数一定要放在形参的最后一个
 */
func sum8(show bool, args ...int) {
}

/*
错误：不定参数放在了第一个位置
 */
/*func sum9(args ...int,show bool)  {
}*/

/*
将不定长参数传递给另外一个不定长参数的函数
 */
func sum10(args ...int) {
	//全部
	sum7(args...)
	//0 - 2
	//sum7(args[:2]...)
	//2 - 最后
	//sum7(args[2:]...)
}
