package dbmap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func DBConfiguer() {
	DB, _ = gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True")

	DB.DB()

	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SingularTable(true)
	DB.LogMode(true)
}
