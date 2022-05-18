package main

import (
	"fmt"
	"gorm/program/dao"
	"gorm/program/models"
	"gorm/program/routers"
)

func main() {

	//1.创建数据库

	//CREATE DATABASE bubble;

	//2.连接数据库
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("Connect database failed: %v\n", err)
		return
	}

	defer dao.DB.Close()

	//3.模型迁移
	dao.DB.AutoMigrate(&models.ToDo{})

	r := routers.SetupRouter()
	r.Run(":8085")
}
