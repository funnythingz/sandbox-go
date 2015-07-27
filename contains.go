package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(Contains(s, "d"))
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
