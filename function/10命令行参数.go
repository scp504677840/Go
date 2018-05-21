package function

import (
	"os"
	"fmt"
)

/*
命令行参数：
go run part10.go
key:0,value:/tmp/go-build313765288/command-line-arguments/_obj/exe/part10

go run part10.go a b
key:0,value:/tmp/go-build313765288/command-line-arguments/_obj/exe/part10
key:1,value:a
key:2,value:b

 */
func main() {
	list := os.Args
	for k, v := range list {
		fmt.Printf("key:%d,value:%s\n", k, v)
	}
}
