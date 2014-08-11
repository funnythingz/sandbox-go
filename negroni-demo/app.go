package main

import (
    "github.com/codegangsta/negroni"
    "net/http"
    "log"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        log.Println(w)
    })

    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":3000")
}
