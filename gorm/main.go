package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //引入gorm
	_ "github.com/jinzhu/gorm/dialects/mysql" //匿名引入mysql驱动，不直接使用，而是使用其内部一些初始化函数
)

//定义一个结构体（在数据库中对应一个表）

type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	//注意书写用户名、密码、主机名称、端口号以及数据库名称
	db, err := gorm.Open("mysql", "root:15willis,@tcp(shiyivei.com:3301)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//操作完之后记得关闭
	defer db.Close()

	//创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	// u1 := UserInfo{ID: 1, Name: "shi", Gender: "male", Hobby: "swimming"}
	// db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //查询表中的第一条数据
	fmt.Printf("u:%#v\n", u)

	//更新,将爱好信息变更
	db.Model(&u).Update("hobby","running")
	
	//删除
	db.Delete(&u)
}
