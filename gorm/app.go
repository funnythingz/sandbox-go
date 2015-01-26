package main

import (
	"./db"
	"github.com/manveru/faker"
	"log"
)

type Person struct {
	Id   int64
	Name string
}

func main() {

	dbmap.DBConfiguer()
	fake, _ := faker.New("en")

	person := Person{Name: fake.Name()}

	log.Println(dbmap.DB.NewRecord(person))
	log.Println(dbmap.DB.Create(&person))
	log.Println(dbmap.DB.NewRecord(person))
	log.Println(dbmap.DB.Save(&person))

	log.Println("first")
	log.Println(dbmap.DB.First(&person))
	log.Println(person)

	var persons []Person
	dbmap.DB.Find(&persons)

	log.Println(persons)

	for k, v := range persons {
		log.Println(k, v.Name)
	}

}
