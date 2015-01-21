package main

import (
	"fmt"
)

type Human struct {
	Name string
}

type Organization struct {
	Members []Human
}

func main() {

	yamada := Human{Name: "yamada"}
	tanaka := Human{Name: "tanaka"}

	humans := []Human{yamada, tanaka}

	fmt.Println(humans)

	organization := Organization{Members: humans}
	fmt.Println(organization)

}
