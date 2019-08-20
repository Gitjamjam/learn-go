package main

import "fmt"

// Goの関数はクロージャ(closure)
// クロージャは自身の外部から変数を参照する関数値
// 参照された変数へアクセスして変えることが可能
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i:= 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
