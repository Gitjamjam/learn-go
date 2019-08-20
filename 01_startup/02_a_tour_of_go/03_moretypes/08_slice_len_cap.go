package main

import "fmt"

func main() {

  // スライスは長さ(length)と容量(capacity)を持っている
  // スライスの長さは要素の数
  // スライスの容量は最初の要素から数えてもととなる配列の要素数

  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)


  s = s[:0]
  printSlice(s)

  s = s[:4]
  printSlice(s)

  s = s[2:]
  printSlice(s)

  // スライスのzセロ値はnull
  var s2 []int
  fmt.Println(s, len(s), cap(s))
  if s == null {
    fmt.Println("null")
  }

}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
