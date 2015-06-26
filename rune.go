package main

import "fmt"

func main() {
	unko := []rune("Hello! うんこだいすき！")
	fmt.Println(string(unko[0:len(unko)]))
}
