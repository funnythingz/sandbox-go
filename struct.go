package main

import(
    "fmt"
)

type T struct {
    Title string
    Body []string
}

func main() {
    t := T{
        "title",
        []string{"body1", "body2"},
    }

    fmt.Println(t)
}
