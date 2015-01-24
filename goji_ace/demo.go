package main

import (
	"fmt"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/", root)
	goji.Get("/:id", hello)
	goji.Serve()
}

func root(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("layout", "content", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"Title": "Welcome"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["id"])
}
