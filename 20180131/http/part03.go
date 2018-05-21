package main

import (
	"os"
	"fmt"
)

func main() {

	file, err := os.Create("./name.txt")

	defer file.Close()

	if err != nil {
		fmt.Println("文件创建失败！！！")
		return
	}

	fmt.Println("文件创建成功！！！")

}
