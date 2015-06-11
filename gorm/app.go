package main

import (
	"./db"
	"./models"
	"github.com/k0kubun/pp"
	"github.com/manveru/faker"
	"log"
)

func main() {

	dbmap.DBConfiguer()

	fake, _ := faker.New("en")

	person := models.Person{Name: fake.Name()}

	dbmap.DB.NewRecord(person)
	dbmap.DB.Create(&person)
	//dbmap.DB.NewRecord(person)
	//dbmap.DB.Save(&person)

	//pp.Println(dbmap.DB.First(&person))
	pp.Println(person)

	var persons []models.Person
	dbmap.DB.Find(&persons)

	pp.Println(persons)

	go func() {
		for k, v := range persons {
			log.Println(k, v.Name)
		}
	}()
}
