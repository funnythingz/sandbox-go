package main

import (
    "fmt"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/acceptlang"
    "net/http"
)

func main() {
    m := martini.Classic()

    m.Get("/", acceptlang.Languages(), func(languages acceptlang.AcceptLanguages) string {
        return fmt.Sprintf("Languages: %s", languages)
    })

    http.ListenAndServe(":8090", m)
}
