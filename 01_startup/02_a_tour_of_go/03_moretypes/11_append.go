package main

import "fmt"

func main() {
  var s []int
  printSlice(s)

  // スライスへ新しい要素を追加するには append を使用
  // 最初の引数は追加元となるスライス
  // 残りの引数は追加する変数群
  s = append(s, 0)
  printSlice(s)

  s = append(s, 1)
  printSlice(s)

  s = append(s, 2, 3, 4)
  printSlice(s)

}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
