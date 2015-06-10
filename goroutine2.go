package main

import (
	"fmt"
	"log"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		log.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		log.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	log.Println("done")
}
