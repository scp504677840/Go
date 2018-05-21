package main

import (
	"errors"
	"fmt"
)

/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
 */
func main() {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("数值不能小于0")
	}
	return f * 10, nil
}
