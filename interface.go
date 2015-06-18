package main

import ()

func main() {

	var animals []interface{}

	animals = append(animals, Dog{Name: "pochi"})
	animals = append(animals, Cat{Name: "mike"})
}

type Animal interface {
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}
