package main

import(
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "time"
)

type Entry struct {
    CreateAt int64
    Message string
}

var DatabaseFile string
var Dbmap *gorp.DbMap

func initDB() {

    DatabaseFile = "test.db"

    db, _ := sql.Open("sqlite3", DatabaseFile)
    Dbmap = &gorp.DbMap{
        Db: db,
        Dialect: gorp.SqliteDialect{},
    }

    Dbmap.AddTableWithName(Entry{}, "entry").SetKeys(true, "CreateAt")

    Dbmap.DropTables()
    Dbmap.CreateTables()

}

func InsertEntry(m string) {
    tx, _ := Dbmap.Begin()
    tx.Insert(&Entry{time.Now().UnixNano(), m})
    tx.Commit()
}

func GetEntryAll() []string {

    list, _ := Dbmap.Select(Entry{}, "select * from entry order by CreateAt desc")

    var result []string

    for _, val := range list {
        entry := val.(*Entry)
        result = append(result, entry.Message)
    }

    return result
}
