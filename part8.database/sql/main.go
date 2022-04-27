package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var db *sql.DB //

const (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "15Willis," //使用上一节中的密码
	database = "master"    //使用默认的master即可
)

func main() {
	//数据库的连接字符串
	connStr := fmt.Sprintf("server=%s;user id =%s; password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connStr) //open函数的两个参数分别是数据库驱动的名称以及连接字符串，返回的第一个值是db
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background() //返回一个非空的Context，不会被取消，没有值也没有截止时间

	err = db.PingContext(ctx) //使用db ping数据库
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected!")

	//log.Println(db == nil)

	one, err := getOne(103)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(one)
}

//func main() {
//	projects := []struct {
//		mascot  string
//		release int
//	}{
//		{"tux", 1991},
//		{"duke", 1996},
//		{"gopher", 2009},
//		{"moby dock", 2013},
//	}
//
//	stmt, err := db.Prepare("INSERT INTO projects(id, mascot, release, category) VALUES( ?, ?, ?, ? )")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
//
//	for id, project := range projects {
//		if _, err := stmt.Exec(id+1, project.mascot, project.release, "open source"); err != nil {
//			log.Fatal(err)
//		}
//	}
//}
