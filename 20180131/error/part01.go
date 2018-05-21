package main

import (
	"fmt"
	"errors"
)

/*
Go 错误处理
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
	Error() string
}

我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
 */
func main() {

	fmt.Println(add(-1, 2))

}

func add(x, y int32) (result int32, err error) {
	if x < 0 || y < 0 {
		return result, errors.New("相加的数不能为0！！！")
	}
	result = x + y
	return result, nil
}
