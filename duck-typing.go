package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	var animals []Animal

	animals = append(animals, &Dog{Name: "pochi"})
	animals = append(animals, &Cat{Name: "mike"})

	for _, animal := range animals {
		pp.Println(animal.GetName())
	}
}

type Animal interface {
	GetName() string
	SetName(name string)
}

type Dog struct {
	Name string
}

func (a *Dog) SetName(name string) {
	a.Name = name
}

func (a *Dog) GetName() string {
	return a.Name
}

type Cat struct {
	Name string
}

func (a *Cat) SetName(name string) {
	a.Name = name
}

func (a *Cat) GetName() string {
	return a.Name
}
