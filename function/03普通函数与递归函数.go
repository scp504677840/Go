package function

import "fmt"

func main() {

	//普通函数，累加求和
	//fmt.Println(plus1())

	//递归函数，累加求和
	fmt.Println(plus2(100))
}

func plus1() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return
}

func plus2(i int) (sum int) {
	if i == 1 {
		return 1
	}
	return i + plus2(i-1)
}
