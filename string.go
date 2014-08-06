package main

import(
    "fmt"
    "os"
)

func main() {
    s := "hello"

    if s[1] != 'e' {
        os.Exit(1)
    }

    fmt.Println(s)

    s = "good bye"

    fmt.Println(s)

    var p *string = &s

    *p = "hoge"

    fmt.Println(*p)
    fmt.Println(s)
}
