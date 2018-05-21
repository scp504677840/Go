package main

import "fmt"

/*
Go 语言结构体
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
Title ：标题
Author ： 作者
Subject：学科
ID：书籍ID

定义结构体
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中一个或多个成员。
type 语句设定了结构体的名称。
结构体的格式如下：
type struct_variable_type struct {
	member definition;
	member definition;
	...
	member definition;
}
一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}

访问结构体成员
如果要访问结构体成员，需要使用点号 (.) 操作符，格式为："结构体.成员名"。
结构体类型变量使用struct关键字定义，实例如下：
 */
func main() {

	var book1 Books
	var book2 Books

	//book1 的描述
	book1.book_id = 123
	book1.author = "Tom"
	book1.subject = "Java"
	book1.title = "Java One"

	//book2 的描述
	book2.book_id = 456
	book2.author = "Jay"
	book2.subject = "Go"
	book2.title = "Go Good"

	/*fmt.Println("书籍ID：", book1.book_id)
	fmt.Println("作者：", book1.author)
	fmt.Println("科目：", book1.subject)
	fmt.Println("标题：", book1.title)

	fmt.Println("书籍ID：", book2.book_id)
	fmt.Println("作者：", book2.author)
	fmt.Println("科目：", book2.subject)
	fmt.Println("标题：", book2.title)*/

	plBook(book1)
	plBook(book2)

	//结构体指针
	var book_ptr *Books
	book_ptr = &book1
	fmt.Println(book_ptr)//book1 的变量地址
	fmt.Println(*book_ptr)//book1 的值

	//使用结构体指针访问结构体成员，使用 "." 操作符：
	fmt.Println("书籍ID：", book_ptr.book_id)
	fmt.Println("作者：", book_ptr.author)
	fmt.Println("科目：", book_ptr.subject)
	fmt.Println("标题：", book_ptr.title)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

/*
结构体作为函数参数
 */
func plBook(books Books)  {
	fmt.Println("书籍ID：", books.book_id)
	fmt.Println("作者：", books.author)
	fmt.Println("科目：", books.subject)
	fmt.Println("标题：", books.title)
}
