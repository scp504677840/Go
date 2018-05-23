package main

import (
	"net"
	"fmt"
)

/*
OSI/RM
应用层
表示层
会话层
传输层
网络层
数据链路层
物理层

TCP/IP
应用层：FTP等等【应用程序根据自定数据格式协议来发送和接收】
传输层：TCP、UDP【源端口号-目的端口号；进程到进程】
网络层：IP（IP地址：逻辑地址）等等（广播风暴）【源IP地址-目的IP地址；主机到主机】
链路层：硬件接口（网卡：物理地址-》MAC地址）等待【源MAC地址-目的MAC地址；设备到设备】

CS模型（C：client S：server）
*/
func main() {
	
	//监听tcp
	listener, err := net.Listen("tcp", ":8080")

	//当监听遇到错误时
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭监听
	defer listener.Close()

	for {

		//监听客户端发送过来的数据
		conn, err := listener.Accept()

		//当监听遇到错误时
		if err != nil {
			fmt.Println(err)
			continue
		}

		//缓存
		buf := make([]byte, 1024)
		//读数据
		length, err := conn.Read(buf)

		//当读数据遇到错误时
		if err != nil {
			fmt.Println(err)
			continue
		}

		//打印结果
		fmt.Println(string(buf[:length]))

		conn.Write([]byte("this is server"))

		//关闭连接
		conn.Close()

	}

}
