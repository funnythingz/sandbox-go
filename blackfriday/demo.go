package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	_ "io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//md, err := ioutil.ReadFile("hoge.md")
	//check(err)
	output := blackfriday.MarkdownCommon([]byte("# hoge"))
	fmt.Println(string(output))
}
