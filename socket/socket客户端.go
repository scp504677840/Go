package main

import (
	"net"
	"fmt"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("this is client"))

	buf := make([]byte, 1024)
	length, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf[:length]))

}
