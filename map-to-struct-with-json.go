package main

import (
	"encoding/json"
	"fmt"
)

type example struct {
	Name    string
	Address string
	Hobbies []string
}

func main() {
	m := make(map[string]interface{})

	m["name"] = "go"
	m["address"] = "japan"
	m["hobbies"] = []string{"watching tv", "listen music"}
	fmt.Println(m)

	var ex example
	MapToStruct(m, &ex)
	fmt.Println(ex)

}

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
