package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//把数据库连接信息写在初始化函数中，这样就不用每次都连接数据库了

func InitMySQL() (err error) {
	dsn := "root:15willis,@tcp(shiyivei.com:3301)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open database failed: %v\n", err)
		return
	}
	//测试连通性
	return DB.DB().Ping()
}
