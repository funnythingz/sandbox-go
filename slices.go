package main

import "fmt"

func main() {

    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }

    fmt.Println("--")

    for num, v := range p {
        fmt.Printf("p[%d] == %d\n", num, v)
    }

}
