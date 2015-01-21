package main

import (
	"fmt"
)

// classみたいなもの
type Body struct {
	B, W, H float64
}

// Methods
func (b *Body) SizeB() float64 {
	return b.B
}

func (b *Body) SizeW() float64 {
	return b.W
}

func (b *Body) SizeH() float64 {
	return b.H
}

// main
func main() {
	b := &Body{90, 55, 70}
	fmt.Println(b, b.SizeB())
}
