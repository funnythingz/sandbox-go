package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	html, err := ioutil.ReadFile("hoge.html")
	check(err)

	p := bluemonday.UGCPolicy()
	fmt.Println(p.Sanitize(string([]byte(html))))
}
