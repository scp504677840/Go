package basis

import "fmt"

/*
for
range
 */
func main() {

	//正常循环
	/*var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	//无限循环
	/*for {
		println("无限循环")
	}*/

	s := "aabbcc"

	for index := range s {
		fmt.Printf("%c",s[index])
	}

	for k, v := range s {
		fmt.Printf("k=%d,v=%c\n", k, v)
	}

}
