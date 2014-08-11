package main

import (
    "net/http"
    _ "log"
    "github.com/codegangsta/negroni"
    "github.com/unrolled/render"
)

type Header struct {
    Title string
    Lead string
}

func main() {

    r := render.New(render.Options{})

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        r.HTML(w, http.StatusOK, "index", &Header{"Hello", "world"})
    })

    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":3000")
}