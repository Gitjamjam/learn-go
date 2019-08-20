package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  // 型変換、明示的な変換が必要
  var f flaot64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)

  // 明示的な型を指定せずに型変換をする場合、
  // 右辺の変数が型を持っていると左辺の変数も同じ型になる
  v := 42
  fmt.Println("v is of type %T\n", v)
}
