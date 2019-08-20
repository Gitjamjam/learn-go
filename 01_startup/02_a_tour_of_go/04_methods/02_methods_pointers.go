package main

import (
  "fmt"
  "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインターレシーバでメソッド宣言
// レシーバの型が、ある型への構文 *Tがあることを意味する
// ポインターレシーバが持つメソッド(Scale)は
// レシーバが指す変数を変更できます
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  v.Scale(10)
  fmt.Println(v.Abs())
}
