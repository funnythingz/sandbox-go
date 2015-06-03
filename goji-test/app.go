package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type User struct {
	Id   int64
	Name string
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	u := User{
		Id:   1,
		Name: c.URLParams["name"],
	}

	j, _ := json.Marshal(u)
	fmt.Fprintf(w, string(j))
}

func Route(m *web.Mux) {
	m.Get("/hello/:name", hello)
}

func main() {
	Route(goji.DefaultMux)
	goji.Serve()
}
