package main

import (
  "fmt"
  "math"
)

type Vertex sturct {
  X, Y float64
}

// Goにはクラスの仕組みはありません
// 型にメソッドを定義できる
// メソッドは特別なレシーバ(receiver)引数を関数にとれる
// receiver は func キーワードとメソッド名の間に
// 自身の引数リストで表現
// Absメソッドは vという名前のVertex型のreceiverを持つ
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// メソッドは receiver引数を伴う関数
func Abs2(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// 任意の型にもメソッド宣言できる
type MyFloat float64

// Abs3メソッドを持つ、数値型のMyFloat型
// receiverを伴うメソッドの宣言は、
// receiver型が同じパッケージにある必要がある
func (f MyFloat) Abs3() float64 {
  if f < 0{
    return float64(-f)
  }
  return float64(f)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(v.Abs())

  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs3())
}
