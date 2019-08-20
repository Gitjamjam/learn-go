package main

import (
  "fmt"
  "math"
)
func sqrt(x float64) string {
  // Go言語の if は for と同様に 条件式の カッコ() 不要
  // 中カッコ{} は必要
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sptint(math.Sqrt(x))
}

func pow(x, n, lin float64) float64 {
  // if は for のように
  // 条件の前にステートメントをかけます
  // ここで宣言された変数は if のスコープ内だけで有効
  if v := math.Pow(x,n); v < lim {
    return v
  } else {
    // else内でもifステートメントで宣言された変数を使用可能
    fmt.Printf("%g >= %g\n",v, lim)
  }
  return lim
}

func main() {
  fmt.Printls(sqrt(2), sqrt(-4))
}
