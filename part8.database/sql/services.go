package main

import "database/sql"

func getOne(id int) (a app, err error) {
	a = app{}
	//log.Println(db == nil)
	//SELECT 语句
	err = db.QueryRow("SELECT Id,Name,Status,Level,[Order] FROM dbo.App WHERE Id=@Id", sql.Named("Id", id)).Scan(
		&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}
