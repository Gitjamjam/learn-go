package main

import "fmt"

func main() {

  // スライスは配列への参照
  // データを格納していなく、元の配列の部分列を指し示している
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Prinln(names)

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)
}
