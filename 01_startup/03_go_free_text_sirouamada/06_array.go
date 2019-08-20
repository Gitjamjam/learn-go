package main

import (
  "fmt"
)

func main() {

  // 配列の定義
  var a [3]int
  b := [2]string{"hoge", "fuga"}

  // 初期値を指定すると「...」表記で長さを省略
  c := [...]int{1, 2}

  // インデックスを指定して値を割り当て
  d := [3]int{1:1,2:10}

}
