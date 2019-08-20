package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe bool = false
  // uint64: 64ビット符号なし整数
  MaxInt uint64 = 1<<64 -1
  // complex128: 複素数型で2つのfloat64型で構成
  z complex128 = cpmlx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("type: %T, Value: %v\n", ToBe, ToBe)
  fmt.Printf("type: %T, Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("type: %T, Value: %v\n", z, z)

  // 変数に初期値を与えず宣言するとゼロ値
  // 数値型：0
  // bool型: false
  // string: ""
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v, %v, %v, %q\n",i, f, b, s)
}
