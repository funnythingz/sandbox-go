package main

import (
	"fmt"
)

func main() {
	values := map[string]string{
		"name": "unko",
		"id":   "20",
	}
	for key, val := range values {
		fmt.Println(key, val)
	}
}
