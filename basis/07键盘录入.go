package basis

import "fmt"

func main() {

	var a int
	fmt.Println("请输入变量：")

	// 情况一：需要输入格式
	//fmt.Scanf("%d",&a)
	//fmt.Println(a)

	// 情况二：不需要输入格式
	fmt.Scan(&a)
	fmt.Println(a)

}
