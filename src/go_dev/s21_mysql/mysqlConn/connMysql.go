package main

import(
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
// 连接mysql
func main(){
	Db, err := sqlx.Open("mysql", "root:123456@tcp(10.143.57.161:3306)/lyydb")
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
	fmt.Println("Connect to mysql success")
	Db.Close()
}

// 增删改查 事务Begin