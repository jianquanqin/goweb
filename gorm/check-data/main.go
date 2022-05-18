package main

import (
	"github.com/jinzhu/gorm"                  //引入gorm
	_ "github.com/jinzhu/gorm/dialects/mysql" //匿名引入mysql驱动，不直接使用，而是使用其内部一些初始化函数
)

type User struct {
	gorm.Model
	Name string
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

	//user1 := User{Name: "shiyivei", Age: 17}
	//db.Create(&user1)
	//
	//user2 := User{Name: "jinzhu", Age: 27}
	//db.Create(&user2)

	//4.查询数据
	var user User
	db.First(&user) //查询第一条数据
	// fmt.Printf("%#v\n", user)

	// var users []User
	// db.Find(&users) //查询所有数据
	// fmt.Printf("%#v\n", users)

	//5.更新
	user.Name = "yivei"
	user.Age = 27
	db.Save(&user) //更新数据
	db.Model(&user).Update("name", "xiaowangzi") //更新数据



	//6.删除
	user.ID = 1
	db.Debug().Delete(&user) //删除数据

}
