package main

import (
	"fmt"
)

type Entity struct {
	Id int
}

func main() {
	var list []interface{}
	hey(&list)
	fmt.Println(list)
}

func hey(list *[]interface{}) {
	hoge := Entity{Id: 1}
	*list = append(*list, hoge)
}
