package main

import (
	"../db"
	"github.com/jinzhu/gorm"
)

type Person struct {
	Id   int64
	Name string
}

var DbMap gorm.DB

func main() {
	dbmap.DBConfiguer()
	DBMigrate()
}

func DBMigrate() {
	dbmap.DB.DropTableIfExists(&Person{})
	dbmap.DB.CreateTable(&Person{})
	dbmap.DB.AutoMigrate(&Person{})

	dbmap.DB.Model(&Person{}).AddIndex("idx_person_name", "name")
}
