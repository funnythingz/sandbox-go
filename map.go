package main

import (
	"fmt"
)

type Human struct {
	Name string
}

func (h *Human) SayName() string {
	return h.Name
}

func main() {
	values := map[string]interface{}{
		"name":  "unko",
		"id":    20,
		"human": Human{Name: "hehey"},
	}
	for key, val := range values {
		fmt.Println(key, val)
	}

	human := values["human"].(Human)
	fmt.Println(human.SayName())
}
