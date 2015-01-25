package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	md, err := ioutil.ReadFile("hoge.md")
	check(err)
	output := blackfriday.MarkdownCommon([]byte(md))
	fmt.Println(string(output))
}
