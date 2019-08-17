package main

import "fmt"

// struct(構造体)はフィールド(field)の集まり
type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1, 2})

  // structのフィールドは ドット. を用いてアクセス
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)

  // structのフィールドは structのポインタを通してアクセスできる
  // フィールド X を持つ structポインタ p があるとき
  // フィールド X にアクセスするには (*p).X のように書ける
  // または省略してp.Xとかける
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
