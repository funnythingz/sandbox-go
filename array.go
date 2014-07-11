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
}
