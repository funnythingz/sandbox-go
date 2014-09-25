package main

import "log"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {

    x, y := split(17)

    log.Println(x)
    log.Println(y)

}
