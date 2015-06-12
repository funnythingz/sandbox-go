package main

import (
	"./db"
	"./models"
	"github.com/k0kubun/pp"
)

func main() {
	dbmap.DBConfiguer()

	person := models.Person{}

	dbmap.DB.First(&person)
	pp.Println(person)
	person.Name = "unko"

	dbmap.DB.Model(&models.Person{}).Update(&person)

	pp.Println(person)
}
