package main

import "fmt"

func sayMessage() func(string) string {

	return func(message string) string {
		return message
	}
}

func main() {

	say := sayMessage()
	fmt.Println(say("test"))
}
