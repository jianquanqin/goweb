package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

const (
	server   = "127.0.0.1"
	port     = 1433
	user     = "shiyivei"
	password = "15Willis,"
	database = "mcr.microsoft.com/mssql/server"
)

var db *sql.DB

func main() {
	connStr := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
		server,
		port,
		user,
		password,
		database,
	)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected")
}
