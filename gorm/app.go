package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Id   int64
	Name string
}

var DbMap gorm.DB

func main() {

	DbInit()

	person := Person{Name: "hoge"}
	log.Println(DbMap.NewRecord(person))
	log.Println(DbMap.Create(&person))
	log.Println(DbMap.NewRecord(person))
	log.Println(DbMap.Save(&person))

	log.Println("first")
	log.Println(DbMap.First(&person))

}

func DbInit() {

	DbMap, _ = gorm.Open("sqlite3", "test.db")

	DbMap.DB()

	DbMap.DB().Ping()
	DbMap.DB().SetMaxIdleConns(10)
	DbMap.DB().SetMaxOpenConns(100)

	DbMap.SingularTable(true)

	Migrate()

}

func Migrate() {
	DbMap.CreateTable(&Person{})
	DbMap.DropTable(&Person{})
	DbMap.DropTableIfExists(&Person{})
	DbMap.AutoMigrate(&Person{})

	DbMap.Model(&Person{}).AddIndex("idx_person_name", "name")
}
