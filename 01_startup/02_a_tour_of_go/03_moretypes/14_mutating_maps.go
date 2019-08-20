package main

import "fmt"

func main() {
  m := make(map[string]int)

  // map の要素への挿入や更新
  m["Answer"] = 42
  // map の要素の取得
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  // map の要素の削除
  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  // map のキーに対する要素が存在するかは2つ目の値で確認
  // 対象キーに対して要素があれば
  // 2つ目の変数はtrue なければ false
  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)

}
