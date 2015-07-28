package main

import (
	"encoding/json"
	"fmt"
)

type Object struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Memo  string `json:"-"`
}

func main() {

	o := Object{Name: "unko", Value: "ばりゅー", Memo: "ここはみえないお"}

	response, _ := json.Marshal(o)
	fmt.Println(string(response))
}
