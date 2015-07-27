package main

import (
	"fmt"
)

func main() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "b", "c", "e"}
	fmt.Println(Difference(s1, s2))
}

func Difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
