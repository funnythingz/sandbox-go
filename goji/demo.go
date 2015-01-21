package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/hello/:id", hello)
	goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["id"])
}
