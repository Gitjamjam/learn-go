package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  // structリテラルは、フィールドの値を列挙して、
  // 新しいstruct初期値を割り当てられる
  v1 = Vertex{1, 2}
  // Name: を使ってフィールドの一部だけを列挙できる
  v2 = Vertex(X: 1)
  v3 = Vertex{}
  // &を先頭につけると新しく割り当てられたstructへ
  // ポインタを戻せる
  p = &Vertex{1,2 }
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
