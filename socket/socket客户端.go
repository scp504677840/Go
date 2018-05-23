package main

import (
	"net"
	"fmt"
)

/*
socket客户端
*/
func main() {

	//主动连接服务器
	conn, err := net.Dial("tcp", ":8080")

	//当连接服务器遇到错误时
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭连接
	defer conn.Close()

	//发送数据
	conn.Write([]byte("this is client"))

	//缓存
	buf := make([]byte, 1024)
	//读取数据
	length, err := conn.Read(buf)

	//当读取数据遇到错误时
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印数据
	fmt.Println(string(buf[:length]))

}
