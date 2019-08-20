package main

import "fmt"

func main() {
  var fnc = f()
  fmt.Println(fnc())
  fmt.Println(fnc())
  fmt.Println(fnc())
}

func f() func() int {
  i := 0
  fmt.Println("start")
  defer fmt.Println("finish")
  panic("パニック")
  fmt.Println("got for")
  return func() int {
    i++
    return i
  }
}
