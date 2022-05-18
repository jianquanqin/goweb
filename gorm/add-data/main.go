package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //引入gorm
	_ "github.com/jinzhu/gorm/dialects/mysql" //匿名引入mysql驱动，不直接使用，而是使用其内部一些初始化函数
)

type User struct {
	ID   int64
	Name *string `gorm:"default:'xiaowangzi'"` //默认值,如过传入空值可以使用指针
	Age  int64
}

func main() {
	//1.连接数据库
	db, err := gorm.Open("mysql", "root:15willis,@tcp(shiyivei.com:3301)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	//show create table users; 查看表的创建语句

	//3.创建数据

	user := User{Name: new(string), Age: 48}
	fmt.Println(db.NewRecord(&user)) //判断主键是否为空
	db.Create(&user)
	fmt.Println(db.NewRecord(&user)) //判断主键是否为空

}
