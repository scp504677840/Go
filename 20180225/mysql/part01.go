package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	//192.168.10.137
	db, err := sql.Open("mysql", "root:scp15335747148@tcp(127.0.0.1:3306)/lab?charset=utf8")
	if err != nil {
		fmt.Println("mysql连接失败！！！", err)
		return
	}
	db.Close()
}
