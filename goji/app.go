package main

import (
	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/", top)
	goji.Get("/:id", hello)
	goji.Serve()
}

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("layout", "content", &ace.Options{DynamicReload: true})

	if err := tpl.Execute(w, map[string]string{"Title": "Welcome"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("layout", "content", &ace.Options{DynamicReload: true})

	if err := tpl.Execute(w, map[string]string{"Title": "Hello, " + c.URLParams["id"] + "!"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
