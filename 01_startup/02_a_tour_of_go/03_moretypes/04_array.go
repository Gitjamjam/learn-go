package main

import "fmt"

func main() {

  // [n]T型は 型Tのn個の変数の配列を表します
  // 以下では string型 2個の配列を宣言
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  // 配列は固定だがスライスは可変
  // 型 []T は 型Tのスライスを表します
  // コロンで区切られた2つのインデックス low と high
  // の境界を指定することでスライスが形成
  primes2 := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  fmt.Printls(s)

}
