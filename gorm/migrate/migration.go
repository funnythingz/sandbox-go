package main

import (
	"../db"
	"../models"
)

func main() {
	dbmap.DBConfiguer()
	DBMigrate()
}

func DBMigrate() {
	dbmap.DB.DropTableIfExists(&model.Person{})
	dbmap.DB.CreateTable(&model.Person{})
	dbmap.DB.AutoMigrate(&model.Person{})

	dbmap.DB.Model(&model.Person{}).AddIndex("idx_person_name", "name")
}
