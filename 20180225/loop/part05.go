package main

import "fmt"

/*
goto label;
..
.
label: statement;
 */
func main() {
	var i = 0
LOOP:
	for ; i < 10; i++ {
		if i == 5 {
			i++
			goto LOOP
		}
		fmt.Println(i)
	}
}
