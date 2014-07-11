package main

import(
    "fmt"
    "strconv"
)

func main() {

    var arrayOfString [3]string

    for i := 0; i < 3; i++ {
        arrayOfString[i] = "hoge" + strconv.Itoa(i)
    }

    fmt.Println(arrayOfString[0])
    fmt.Println(arrayOfString)

    for i := 0; i < len(arrayOfString); i++ {
        fmt.Println(arrayOfString[i])
    }

    // forEachみたいな感じ
    for i, v := range arrayOfString {
        fmt.Println(i)
        fmt.Println(v)
    }
}
