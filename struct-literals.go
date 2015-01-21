package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
	t = Vertex{Y: 3}  // X:0 and Y:3
)

func main() {
	fmt.Println(p, q, r, s, t)
}
