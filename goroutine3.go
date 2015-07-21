package main

import (
	"fmt"
	"time"
)

func main() {
	task := make(chan string)
	taskquit := make(chan bool)
	workerquit := make(chan bool)

	go func() {
	loop:
		for {
			select {
			case <-taskquit:
				workerquit <- true
				break loop
			case job := <-task:
				fmt.Println(job)
			}
		}
	}()

	go func() {
		for n := 0; n < 3; n++ {
			task <- fmt.Sprintf("お仕事%03d", n+1)
			time.Sleep(1 * time.Second)
		}
		taskquit <- true
	}()

	<-workerquit
}
