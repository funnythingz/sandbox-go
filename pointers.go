package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    p := Vertex{1, 2}
    fmt.Println(p)

    q := &p
    q.X = 3
    fmt.Println(p)
}
