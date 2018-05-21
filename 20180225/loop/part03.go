package main

import "fmt"

/*
for { }
 */
func main() {

	for {
		fmt.Println("死循环")
	}

	for true {
		fmt.Println("死循环")
	}
}
