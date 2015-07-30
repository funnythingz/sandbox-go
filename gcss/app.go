package main

import (
	"fmt"
	"github.com/yosssi/gcss"
)

func main() {
	cssPath, err := gcss.CompileFile("styles.gcss")
	if err != nil {
		panic(err)
	}
	fmt.Println(cssPath)
}
