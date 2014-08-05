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
  var keitan string
  keitan = "良い日だね！けいたん"
  fmt.Println(keitan)
}

func smoothLoop() {
  smooth := []string {"keitan", "ipe", "hrtn?"}
  for i := 0; i < len(smooth); i++ {
    fmt.Println(smooth[i])
  }
}
