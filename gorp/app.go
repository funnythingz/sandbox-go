package main

import(
    "fmt"
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
)

type Person struct {
    Id int
    Name string
}

func main() {

    DatabaseFile := "test.db"

    db, _ := sql.Open("sqlite3", DatabaseFile)
    dbmap := &gorp.DbMap{
        Db: db,
        Dialect: gorp.SqliteDialect{},
    }

    dbmap.AddTableWithName(Person{}, "person").SetKeys(true, "Id")

    dbmap.DropTables()
    dbmap.CreateTables()

    tx, _ := dbmap.Begin()
    for i := 0; i < 10; i++ {
        tx.Insert(&Person{
            i, fmt.Sprintf("hoge%03d", i),
        })
    }
    tx.Commit()

    list, _ := dbmap.Select(Person{}, "select * from person")
    for _, val := range list {
        person := val.(*Person)
        fmt.Printf("%d, %s\n", person.Id, person.Name)
    }

    obj, _ := dbmap.Get(Person{}, 1)
    person := obj.(*Person)

    fmt.Println(person.Name)

}
