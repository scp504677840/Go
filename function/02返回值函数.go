package function

import "fmt"

func main() {
	fmt.Println(sum6(1, 2, true))
}

/*
写法一：无命名返回值
 */
func add1(a, b int) int {
	return a + b
}

/*
写法二：有命名返回值【推荐】
*/
func add2(a, b int) (result int) {
	result = a + b
	return result
}

/*
写法三：多返回值，无命名
 */
func add3() (int, bool) {
	return 99, true
}

/*
写法四：多返回值，有命名【推荐】
 */
func add4() (result int, show bool) {
	result = 10
	show = false
	return result, show
}
