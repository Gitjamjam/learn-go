package main

import "fmt"

// 関数は0個以上の引数を取ることが可能
// int型の2つの引数
// 変数の後ろに型名
func add(x int, y int) int {
  return x + y
}

// 関数の2つ以上の引数が同じ型なら、最後にまとめることが可能
func add02(x, y int) int {
  return x + y
}

// 関数は複数の戻り値を返すことができる
// ここでは2つのstringを返す
func swap(x, y string) (string, string) {
  return y,x
}

// naked return
// 関数の戻り値となる変数に名前をつけられる
// 名前をつけた戻り値の変数を使うと return だけで戻せる
func split(sum int)(x, y int) {
  x = sum * 4 /9
  y = sum - x
  return
}

func main() {
  fmt.Println(add(42, 13))

  a,b := swap("hello", "world")
  fmt.Println(a,b)

  fmt.Println(split(17))
}
