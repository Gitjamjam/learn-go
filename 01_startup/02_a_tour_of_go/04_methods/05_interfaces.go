package main

import (
  "fmt"
  "math"
)

// インターフェース型はメソッドの
// シグニチャの集まり
type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f
  a = &v

  // AbsメソッドがVertexではなく*Vertexの定義であり
  // VertexがAbserインターフェースを実装していないため
  // エラーになる
  a = v
  fmt.Println(a.Abs())
}
