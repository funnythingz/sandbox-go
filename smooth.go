package main

import "fmt"

func main() {
    helloKeitan()
    goodDayKeitan()
    smoothLoop()
}

func helloKeitan() {
    keitan := "こんにちは！けいたん"
    fmt.Println(keitan)
}

func goodDayKeitan() {
    keitan := "良い日だね！けいたん"
    fmt.Println(keitan)
}

func smoothLoop() {

    smooth := []string {"keitan", "ipe", "hrtn?"}

    for _, v := range smooth {
        fmt.Println(v)
    }

}
