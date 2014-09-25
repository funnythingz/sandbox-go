package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
    v := new(Vertex)
    fmt.Println(v)
    v.X, v.Y = 11, 9
    fmt.Println(v)

    r := new(Vertex)
    r.X, r.Y = 1, 2
    fmt.Println(r)
}
