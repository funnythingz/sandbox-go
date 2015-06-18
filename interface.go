package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	var animals []interface{}

	animals = append(animals, Dog{Name: "pochi"})
	animals = append(animals, Cat{Name: "mike"})

	pp.Println(animals)
}

type Animal interface {
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}
