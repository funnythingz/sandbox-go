package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/acceptlang"
)

func main() {
	m := martini.Classic()

	m.Get("/", acceptlang.Languages(), func(languages acceptlang.AcceptLanguages) string {
		return fmt.Sprintf("Languages: %s", languages)
	})

	http.ListenAndServe(":3000", m)
}
