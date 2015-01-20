package main

import (
	"log"
)

type T struct {
	Title string
	Body  []string
}

type Vertex struct {
	X int
	Y int
}

func main() {
	t := T{
		"title",
		[]string{"body1", "body2"},
	}
	log.Println(t)

	v := Vertex{1, 2}
	log.Println(v)
	log.Println(v.X)
	log.Println(v.Y)
}
