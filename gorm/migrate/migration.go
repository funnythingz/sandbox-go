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
	//dbmap.DB.DropTableIfExists(&models.Person{})
	//dbmap.DB.CreateTable(&models.Person{})
	dbmap.DB.AutoMigrate(&models.Person{})
}
