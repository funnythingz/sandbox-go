package main

import "log"

func sayMessage() func(string) string {

    return func(message string) string {
        return message
    }
}

func main() {

    say := sayMessage()
    log.Println(say("test"))
}
