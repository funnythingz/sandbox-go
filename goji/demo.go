package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/", welcome)
	goji.Get("/:id", hello)
	goji.Serve()
}

func welcome(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["id"])
}
