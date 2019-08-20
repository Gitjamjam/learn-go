package main

import (
  "fmt"
  "math"
)

// 型にメソッドを実装していくことで
// インターフェースを実装
// インターフェースを実装することで明示的宣言は必要ない
type I interface {
  M()
}

type T struct {
  S string
}

func (t T) M() {
  fmt.Println(t.S)
}

func (t *T) M2() {
  fmt.Println(t.S)
}

type F float64

func main() {
  var i I = T{"hello"}
  i.M()
}
