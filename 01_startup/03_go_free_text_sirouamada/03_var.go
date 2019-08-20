package main

import (
  "fmt"
)
var a int
var b int = 1
var c = 0

var (
  d int
  e int = 1
  f = 0
)
var a1, a2 int
var b1, b2 int = 0, 1
var c1, c2 = 0, 1

const zzz string = "xyz"
const c4 = iota

func main() {
  g := 1
  h, i := "1=1", "2=2"
  fmt.Println("g=",g,"h=",h,"i=",i)

  for j := 0; j < 10; j++ {
    fmt.Println(c4)
  }
}
