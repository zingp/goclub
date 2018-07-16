package main

import(
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"math/rand"
	"time"
)

type UserInfo struct{
	Id int    		`db:"id"`
	Name string		`db:"name"`
	Sex string      `db:"sex"`
	Age int         `db:"age"`
	Email string    `db:"email"`
}

// 连接mysql
func main(){
	Db, err := sqlx.Open("mysql", "root:123456@tcp(10.143.57.161:3306)/lyydb")
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
	fmt.Println("Connect to mysql success")
	defer Db.Close()
	
	start := time.Now().UnixNano()
	// create(Db)
	// insert(Db)
	// selectDb(Db)
	eventDb(Db)
	end := time.Now().UnixNano()
	fmt.Printf("cost time=%d ms", (start - end)/1000/1000)
}

// 增删改查 事务Begin

func create(Db *sqlx.DB){
	createSql := `create table user_info (
		id int primary key auto_increment,
		name varchar(50),
		sex enum('male', 'female'),
		age int(3),
		email varchar(64)
		);`
	result, err := Db.Exec(createSql)
	if err != nil {
		fmt.Printf("create table error:%v", err)
		return
	}

	fmt.Println("create result:", result)
}

func insert(Db *sqlx.DB) {
	for i:=0; i<1000000;i++ {
		name := fmt.Sprintf("user%d", i)
		sex := "male"
		if i % 2 == 1 {
			sex = "female"
		}
		age := rand.Intn(100)
		email := fmt.Sprintf("%d@qq.com", rand.Int63())

		// fmt.Println(name, sex, age, email)

		_, err := Db.Exec("insert into user_info(name, sex, age, email) values(?,?,?,?)",
		name, sex, age, email)
		if err != nil {
			fmt.Println("insert error:", err)
			return
		}
		
		/*
		// 成功后获取id
		result, err := Db.Exec("insert into user_info(name, sex, age, email) values(?,?,?,?)",
		name, sex, age, email)
		if err != nil {
			fmt.Println("insert error:", err)
			return
		}
		user_id := result.LastInsertId()
		*/
	}
}


func selectDb(Db *sqlx.DB) {
	// 查询单行
	var userInfo UserInfo
	err := Db.Get(&userInfo, "select id,name,sex,age,email from user_info where id=?", 1)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	fmt.Println("userInfo::", userInfo)

	// 查询多行
	var userInfoList []*UserInfo
	err = Db.Select(&userInfoList, "select id,name,sex,age,email from user_info where id<?", 5)
	if err != nil {
		fmt.Println("select error:", err)
		return
	}
	fmt.Println("userInfoList::", userInfoList)
}


func eventDb(Db *sqlx.DB) {
	conn, _ := Db.Begin()    // 使用事务
	_, err := Db.Exec("insert into user_info(name, sex, age, email) values(?,?,?,?)",
	"nameXX", "male",55,"666@sg.com")
	if err != nil {
		conn.Rollback()   // 事务回滚
		fmt.Println("insert1 error:", err)
		return
	}

	_, err = Db.Exec("insert into user_info(name, sex, age, email) values(?,?,?,?)",
	"nameYY", "male", 25.5, "777@sg.com")
	if err != nil {
		conn.Rollback()    // 事务回滚
		fmt.Println("insert2 error:", err)
		return
	}
	
	conn.Commit()    // 事务提交
	fmt.Print("commit success!")

	// 查询
	var userInfo UserInfo
	err = Db.Get(&userInfo, "select id,name,sex,age,email from user_info where email=?", "666@sg.com")
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	fmt.Println("userInfo::", userInfo)
}