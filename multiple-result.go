package main

import(
        "fmt"
        "log"
      )

func swap(x, y string) (string, string) {
    return y, x
}

func magic(x, y int) (int, int) {
    return (x + y), (x * y)
}

func main() {

    a, b := swap("hello", "world")
    fmt.Println(a, b)

    x, y := magic(2, 5)
    log.Println(x, y)

}
