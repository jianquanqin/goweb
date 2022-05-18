package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"                  //引入gorm
	_ "github.com/jinzhu/gorm/dialects/mysql" //匿名引入mysql驱动，不直接使用，而是使用其内部一些初始化函数
	"time"
)

type User struct {
	gorm.Model   //匿名结构体，继承gorm.Model
	Name         string
	Age          sql.NullInt64
	BirthDate    *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号，唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置自增长
	Address      string  `gorm:index:"addr"`      //给Address字段设置addr索引
	IgnoreMe     int     `gorm:"-"`               //忽略该字段

}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func (Animal) FixTableName() string {
	return "Animal"
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
	db.SingularTable(true) //禁用复数

	//创建表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	db.Table("xiaodongwu").CreateTable(&User{})
}
