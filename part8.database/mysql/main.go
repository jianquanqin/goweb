package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		log.Println(err)
		return
	}

	fmt.Println("connected")

}
