package basis

import "fmt"

/*
字符
 */
func main() {

	var ch byte
	ch = 97
	fmt.Printf("%c %d\n", ch, ch)

	ch = 'a'
	fmt.Printf("%c %d\n", ch, ch)

	//字符的数值形式
	fmt.Printf("%d %d\n", 'A', 'a')

	//大小写转换
	fmt.Printf("%c\n", 'a'-32)
	fmt.Printf("%c\n", 'A'+32)

	//字符串
	var s string = "hi"
	fmt.Println(s)
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[1])
	fmt.Println(len(s))
}
