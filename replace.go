package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := `
    Hello
    World     
    World
    World
    World
    `
	fmt.Println(str)
	fmt.Println(strings.Replace(str, "\n", "", -1))

	r := regexp.MustCompile(`([\s]{2,}|\n)`)
	fmt.Println(r.ReplaceAllString(str, " "))
}
