package main

import (
        "log"
        "time"
       )

func main() {

    log.Print("started")

    hogeDone := make(chan bool)
    fooDone := make(chan bool)

    go func() {
        log.Print("hoge start")
        time.Sleep(1 * time.Second)
        log.Print("hoge done")
        hogeDone <- true
    }()

    go func() {
        log.Print("foo start")
        time.Sleep(2 * time.Second)
        log.Print("foo done")
        fooDone <- true
    }()

    <- hogeDone
    <- fooDone

    log.Print("done")

}
