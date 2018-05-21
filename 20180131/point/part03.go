package main

import "fmt"

/*
Go 语言指针数组
 */
func main() {

	//声明一个整型指针数组
	var ptr [3]*int
	//ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：

	nums := []int{10, 100, 200}

	for i := 0; i < len(nums); i++ {
		//整型地址赋给指针数组
		ptr[i] = &nums[i]
	}

	for i := 0; i < len(ptr); i++ {
		fmt.Println(*ptr[i])
	}

	//千万不能用下面这种方式取数组中每一个值的地址，否则取的地址都是v的地址
	/*for k, v := range nums {
		fmt.Println(k,&v)
		ptr[k] = &v
	}

	for k, v := range ptr {
		fmt.Println(k,*v)
	}*/

}
